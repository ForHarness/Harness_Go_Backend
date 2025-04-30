// package main

// import (
// 	"database/sql"
// 	_ "github.com/go-sql-driver/mysql"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"todo_app/handlers"
// 	"todo_app/services"

// 	"github.com/rs/cors"
// )

// func main() {
// 	// Initialize database connection
// 	dsn := "avnadmin:AVNS_j9dszxAT3x87Lf9xOKb@tcp(mysql-bafa257-shubham1997.g.aivencloud.com:20602)/todo_app"
// 	db, err := sql.Open("mysql", dsn)
// 	if err != nil {
// 		log.Fatal("Error connecting to the database:", err)
// 	}
// 	defer db.Close()

// 	// Initialize TodoService
// 	todoService := &services.TodoServiceImpl{DB: db}

// 	// Initialize TodoHandler with TodoService
// 	todoHandler := handlers.NewTodoHandler(todoService)

// 	// Define routes
// 	mux := http.NewServeMux()

// 	mux.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
// 		switch r.Method {
// 		case http.MethodGet:
// 			todoHandler.GetAllTodos(w, r)
// 		case http.MethodPost:
// 			todoHandler.CreateTodo(w, r)
// 		default:
// 			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 		}
// 	})

// 	mux.HandleFunc("/todos/", func(w http.ResponseWriter, r *http.Request) {
// 		if r.Method == http.MethodGet {
// 			todoHandler.GetTodoByID(w, r)
// 		} else if r.Method == http.MethodPut {
// 			todoHandler.UpdateTodo(w, r)
// 		} else if r.Method == http.MethodDelete {
// 			todoHandler.DeleteTodo(w, r)
// 		} else {
// 			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 		}
// 	})

// 	// CORS middleware
// 	corsHandler := cors.Default().Handler(mux)

// 	// Start the server
// 	port := ":8080"
// 	fmt.Println("Server started on port", port)
// 	log.Fatal(http.ListenAndServe(port, corsHandler))
// }

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
	dsn := "avnadmin:AVNS_kWWFwhOYAkcQb_eUmB7@tcp(mysql-35704c68-tododata.h.aivencloud.com:28989)/todo_app"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	// Close DB only if connection is successful
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
	fmt.Println("Server started on port", port)
	log.Fatal(http.ListenAndServe(port, handler))
}

