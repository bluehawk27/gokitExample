package service

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/bluehawk27/gokitExample/todo/pkg/db"

	"github.com/bluehawk27/gokitExample/todo/pkg/io"
)

// TodoService describes the service.
type TodoService interface {
	Get(ctx context.Context) (t []io.Todo, error error)
	Add(ctx context.Context, todo io.Todo) (t io.Todo, error error)
	SetComplete(ctx context.Context, id string) (error error)
	RemoveComplete(ctx context.Context, id string) (error error)
	Delete(ctx context.Context, id string) (error error)
}

var s = db.NewStore()

type basicTodoService struct{}

func (b *basicTodoService) Get(ctx context.Context) ([]io.Todo, error) {
	t, err := s.List(ctx)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (b *basicTodoService) Add(ctx context.Context, todo io.Todo) (io.Todo, error) {
	fmt.Println("trying to add todo: ", todo)
	t, err := s.Add(ctx, todo)
	return *t, err
}

func (b *basicTodoService) SetComplete(ctx context.Context, id string) error {
	idInt, intErr := strconv.ParseInt(id, 10, 32)
	if intErr != nil {
		return errors.New("error parsing ID:" + id)
	}

	err := s.CompleteTodo(ctx, true, idInt)
	return err
}

func (b *basicTodoService) RemoveComplete(ctx context.Context, id string) error {
	idInt, intErr := strconv.ParseInt(id, 10, 32)
	if intErr != nil {
		return errors.New("error parsing ID:" + id)
	}

	err := s.CompleteTodo(ctx, false, idInt)
	return err
}

func (b *basicTodoService) Delete(ctx context.Context, id string) error {
	fmt.Println("Im here Deleting In the service:", id)
	idInt, intErr := strconv.ParseInt(id, 10, 32)
	if intErr != nil {
		return errors.New("error parsing ID:" + id)
	}

	err := s.DeleteTodo(ctx, idInt)
	return err
}

// NewBasicTodoService returns a naive, stateless implementation of TodoService.
func NewBasicTodoService() TodoService {
	return &basicTodoService{}
}

// New returns a TodoService with all of the expected middleware wired in.
func New(middleware []Middleware) TodoService {
	var svc TodoService = NewBasicTodoService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
