// services/todo_service.go
package services

import (
    "database/sql"
    "strconv"
    "todo_app/models"
)

// TodoService defines methods for interacting with todos
type TodoService interface {
    GetAllTodos() ([]models.Todo, error)
    GetTodoByID(id string) (models.Todo, error)
    CreateTodo(todo models.Todo) (models.Todo, error)
    UpdateTodo(id string, todo models.Todo) (models.Todo, error)
    DeleteTodo(id string) error
}

// TodoServiceImpl implements the TodoService interface
type TodoServiceImpl struct {
    DB *sql.DB // Database connection
}

// GetAllTodos returns all todos
func (ts *TodoServiceImpl) GetAllTodos() ([]models.Todo, error) {
    query := "SELECT id, title, description, completed FROM todos"
    rows, err := ts.DB.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var todos []models.Todo
    for rows.Next() {
        var todo models.Todo
        if err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed); err != nil {
            return nil, err
        }
        todos = append(todos, todo)
    }
    return todos, nil
}

// GetTodoByID returns a todo by ID
func (ts *TodoServiceImpl) GetTodoByID(id string) (models.Todo, error) {
    query := "SELECT id, title, description, completed FROM todos WHERE id = ?"
    var todo models.Todo
    err := ts.DB.QueryRow(query, id).Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed)
    if err != nil {
        return models.Todo{}, err
    }
    return todo, nil
}

// CreateTodo creates a new todo
func (ts *TodoServiceImpl) CreateTodo(todo models.Todo) (models.Todo, error) {
    query := "INSERT INTO todos (title, description, completed) VALUES (?, ?, ?)"
    result, err := ts.DB.Exec(query, todo.Title, todo.Description, todo.Completed)
    if err != nil {
        return models.Todo{}, err
    }
    id, err := result.LastInsertId()
    if err != nil {
        return models.Todo{}, err
    }
    todo.ID = int(id)
    return todo, nil
}


// UpdateTodo updates an existing todo
func (ts *TodoServiceImpl) UpdateTodo(id string, updatedTodo models.Todo) (models.Todo, error) {
    query := "UPDATE todos SET title = ?, description = ?, completed = ? WHERE id = ?"
    _, err := ts.DB.Exec(query, updatedTodo.Title, updatedTodo.Description, updatedTodo.Completed, id)
    if err != nil {
        return models.Todo{}, err
    }
    updatedTodo.ID, _ = strconv.Atoi(id)
    return updatedTodo, nil
}

// DeleteTodo deletes a todo by ID
func (ts *TodoServiceImpl) DeleteTodo(id string) error {
    query := "DELETE FROM todos WHERE id = ?"
    _, err := ts.DB.Exec(query, id)
    if err != nil {
        return err
    }
    return nil
}
