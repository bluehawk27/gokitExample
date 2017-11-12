package db

import (
	"context"
	"fmt"

	"github.com/bluehawk27/gokitExample/todo/pkg/io"
	_ "github.com/go-sql-driver/mysql" // Needed for sqlx.
	"github.com/jmoiron/sqlx"
)

type StoreInterface interface {
	List(ctx context.Context) ([]io.Todo, error)
	GetTodoByID(ctx context.Context, id int64) (*io.Todo, error)
	Add(ctx context.Context, todo io.Todo) (*io.Todo, error)
	UpdateTodo(ctx context.Context, todo io.Todo, id int64) (*io.Todo, error)
	CompleteTodo(ctx context.Context, complete bool, id int64) error
	DeleteTodo(ctx context.Context, id int64) error
}

type store struct {
	db *sqlx.DB
}

func NewStore() StoreInterface {
	s := &store{}
	s.connect()
	return s
}

func (s *store) Ping() bool {
	return s.db.Ping() == nil
}

// Connect connects to the database
func (s *store) connect() error {
	if s.db != nil {
		return nil
	}
	var err error
	s.db, err = sqlx.Connect("mysql", "root:@tcp(127.0.0.1:3306)/todo_app?parseTime=true")
	if err != nil {
		return err
	}

	return nil
}

func (s *store) List(ctx context.Context) ([]io.Todo, error) {
	todos := []io.Todo{}

	err := s.db.Select(&todos, getTodos)
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (s *store) GetTodoByID(ctx context.Context, id int64) (*io.Todo, error) {
	var todo io.Todo
	err := s.db.Get(&todo, getTodoByID, id)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func (s *store) Add(ctx context.Context, todo io.Todo) (*io.Todo, error) {
	tx := s.db.MustBegin()
	res, err := tx.Exec(addTodo, todo.Title, todo.Complete)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	id, _ := res.LastInsertId()

	return s.GetTodoByID(ctx, id)
}

func (s *store) UpdateTodo(ctx context.Context, todo io.Todo, id int64) (*io.Todo, error) {
	_, err := s.db.Exec(updateTodo, todo.Title, todo.Complete, id)
	if err != nil {
		return nil, err
	}

	return s.GetTodoByID(ctx, id)
}

func (s *store) CompleteTodo(ctx context.Context, complete bool, id int64) error {
	_, err := s.db.Exec(setComplete, complete, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *store) DeleteTodo(ctx context.Context, id int64) error {
	fmt.Println("im here Deleting")
	tx := s.db.MustBegin()
	_, err := s.db.Exec(deleteTodo, id)
	if err != nil {
		tx.Rollback()
		fmt.Println(err)
		return err
	}
	err = tx.Commit()
	return err

}
