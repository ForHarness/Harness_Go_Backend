package handlers

import (
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
    "todo_app/models"
)

// Define a mock TodoService implementation for testing
type mockTodoService struct{}

func (mts *mockTodoService) GetAllTodos() ([]models.Todo, error) {
    // Implement the mock GetAllTodos method for testing
    todos := []models.Todo{
        {ID: 1, Title: "Example Todo 1", Description: "Example Description 1", Completed: false},
        {ID: 2, Title: "Example Todo 2", Description: "Example Description 2", Completed: true},
    }
    return todos, nil
}

func (mts *mockTodoService) GetTodoByID(id string) (models.Todo, error) {
    // Implement the mock GetTodoByID method for testing
    todo := models.Todo{ID: 1, Title: "Example Todo", Description: "Example Description", Completed: false}
    return todo, nil
}

func (mts *mockTodoService) CreateTodo(todo models.Todo) (models.Todo, error) {
    // Implement the mock CreateTodo method for testing
    return todo, nil
}

func (mts *mockTodoService) UpdateTodo(id string, updatedTodo models.Todo) (models.Todo, error) {
    // Implement the mock UpdateTodo method for testing
    return updatedTodo, nil
}

func (mts *mockTodoService) DeleteTodo(id string) error {
    // Implement the mock DeleteTodo method for testing
    return nil
}

func TestGetAllTodos(t *testing.T) {
    // Create a mock TodoService instance
    todoService := &mockTodoService{}

    // Create a TodoHandler instance with the mock TodoService
    todoHandler := NewTodoHandler(todoService)

    // Create a new HTTP request for the GET /todos endpoint
    req, err := http.NewRequest("GET", "/todos", nil)
    if err != nil {
        t.Fatal(err)
    }

    // Create a response recorder to record the response
    rr := httptest.NewRecorder()

    // Serve the HTTP request to the recorder
    handler := http.HandlerFunc(todoHandler.GetAllTodos)
    handler.ServeHTTP(rr, req)

    // Check the status code
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    // Check the response body
    var todos []models.Todo
    if err := json.NewDecoder(rr.Body).Decode(&todos); err != nil {
        t.Errorf("failed to decode response body: %v", err)
    }

    // Check the length of todos
    if len(todos) != 2 {
        t.Errorf("expected 2 todos, got %d", len(todos))
    }

    // Check the content of the todos
    expectedTodos := []models.Todo{
        {ID: 1, Title: "Example Todo 1", Description: "Example Description 1", Completed: false},
        {ID: 2, Title: "Example Todo 2", Description: "Example Description 2", Completed: true},
    }
    for i, todo := range todos {
        if todo != expectedTodos[i] {
            t.Errorf("unexpected todo at index %d: got %v, want %v", i, todo, expectedTodos[i])
        }
    }
}
