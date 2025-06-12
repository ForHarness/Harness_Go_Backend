package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"todo_app/handlers"
	"todo_app/services"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/cors"
)

func main() {
	// Initialize database connection
	dsn := "avnadmin:AVNS_GaiGgz2LtGMCFMVy9Qu-04@tcp(mysql-3e7e5281-datatodo.c.aivencloud.com:13540)/todo_app?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("❌ Error opening database: %v", err)
	}

	// Ping the database to verify connection
	if err := db.Ping(); err != nil {
		log.Fatalf("❌ Error pinging database: %v", err)
	}

	log.Println("✅ Successfully connected to database")
	defer db.Close()

	// Initialize TodoService
	todoService := &services.TodoServiceImpl{DB: db}

	// Initialize TodoHandler with TodoService
	todoHandler := handlers.NewTodoHandler(todoService)

	// Define routes
	mux := http.NewServeMux()

	mux.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			todoHandler.GetAllTodos(w, r)
		case http.MethodPost:
			todoHandler.CreateTodo(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/todos/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			todoHandler.GetTodoByID(w, r)
		case http.MethodPut:
			todoHandler.UpdateTodo(w, r)
		case http.MethodDelete:
			todoHandler.DeleteTodo(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// CORS middleware with custom settings
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Frontend ELB
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})

	// Wrap the router with CORS
	handler := c.Handler(mux)

	// Start the server
	port := "0.0.0.0:8080"
	fmt.Println("🚀 Server started on port", port)
	log.Fatal(http.ListenAndServe(port, handler))
}
