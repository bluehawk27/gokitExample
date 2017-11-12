package db

const getTodos = "SELECT id, title, complete FROM todo"

const getTodoByID = `SELECT id, title, complete FROM todo WHERE id = ?`

const addTodo = `INSERT INTO todo title, complete VALUES (?, ?)`

const updateTodo = `UPDATE todo SET title = ?, complete = ? WHERE id = ?`

const deleteTodo = `DELETE FROM todo WHERE id = ?`

const setComplete = `UPDATE todo SET complete = ? WHERE id = ?`
