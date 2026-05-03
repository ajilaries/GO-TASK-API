package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"go-task-api/internal/handlers"
	"go-task-api/internal/middleware"
	"go-task-api/pkg/database"
)

func main() {

	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	// Connect DB
	database.ConnectDB()

	// Router
	router := mux.NewRouter()

	// 🌐 Global Middleware
	router.Use(middleware.Logging)
	router.Use(middleware.CORS)         

	// 🔓 Public Routes
	router.HandleFunc("/register", handlers.Register).Methods("POST")
	router.HandleFunc("/login", handlers.Login).Methods("POST")

	// 🔒 Protected Routes
	protected := router.PathPrefix("/tasks").Subrouter()
	protected.Use(middleware.AuthMiddleware)

	// protected.HandleFunc("", handlers.GetTasks).Methods("GET")
	protected.HandleFunc("", handlers.CreateTask).Methods("POST")
	protected.HandleFunc("/{id}", handlers.GetTaskByID).Methods("GET")
	protected.HandleFunc("/{id}", handlers.UpdateTask).Methods("PUT")
	protected.HandleFunc("/{id}", handlers.DeleteTask).Methods("DELETE")

	log.Println("🚀 Server running on :8080")
	http.ListenAndServe(":8080", router)
}