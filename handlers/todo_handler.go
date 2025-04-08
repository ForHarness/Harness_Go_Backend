// handlers/todo_handler.go
package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"todo_app/models"
	"todo_app/services"
)

// TodoHandler handles requests related to todos
type TodoHandler struct {
	// Define any dependencies or fields needed for the handler
	TodoService services.TodoService
}

// NewTodoHandler creates a new TodoHandler
func NewTodoHandler(todoService services.TodoService) *TodoHandler {
	return &TodoHandler{TodoService: todoService}
}

// GetAllTodos returns all todos
func (th *TodoHandler) GetAllTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := th.TodoService.GetAllTodos()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching todos: %s", err.Error()), http.StatusInternalServerError)
		return
	}
	respondWithJSON(w, http.StatusOK, todos)
}

// GetTodoByID returns a todo by ID
func (th *TodoHandler) GetTodoByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID := vars["id"]

	todo, err := th.TodoService.GetTodoByID(todoID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching todo: %s", err.Error()), http.StatusInternalServerError)
		return
	}
	respondWithJSON(w, http.StatusOK, todo)
}

// CreateTodo creates a new todo
func (th *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, fmt.Sprintf("Error decoding request body: %s", err.Error()), http.StatusBadRequest)
		return
	}

	newTodo, err := th.TodoService.CreateTodo(todo)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating todo: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusCreated, newTodo)
}

// UpdateTodo updates an existing todo
func (th *TodoHandler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID := vars["id"]

	var updatedTodo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&updatedTodo); err != nil {
		http.Error(w, fmt.Sprintf("Error decoding request body: %s", err.Error()), http.StatusBadRequest)
		return
	}

	updatedTodo, err := th.TodoService.UpdateTodo(todoID, updatedTodo)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error updating todo: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, updatedTodo)
}

// DeleteTodo deletes a todo by ID
func (th *TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID := vars["id"]

	err := th.TodoService.DeleteTodo(todoID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error deleting todo: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Todo deleted successfully"})
}

func respondWithJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
