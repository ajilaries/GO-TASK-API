package main

import (
	"log"
	"net/http"

	"go-task-api/internal/handlers"
	"go-task-api/internal/middleware"
	"go-task-api/pkg/database"
	"github.com/joho/godotenv"
	"github.com/go-chi/chi/v5"
	
)

func main() {
	//load .env file
	err:= godotenv.Load()
	if err!=nil{
		log.Println("No .env file found")
	}
	//db connectivity
	database.ConnectDB()

	r := chi.NewRouter()

	// 🌐 Global middleware
	r.Use(middleware.CORS)
	r.Use(middleware.Logging)

	// 🔓 Public routes
	r.Post("/register", handlers.Register)
	r.Post("/login", handlers.Login)

	// 🔒 Protected routes
	r.Route("/tasks", func(r chi.Router) {

		r.Use(middleware.AuthMiddleware)

		r.Get("/", handlers.GetTasks)
		r.Post("/", handlers.CreateTask)
		r.Put("/{id}", handlers.UpdateTask)
		r.Delete("/{id}", handlers.DeleteTask)
		r.Get("/{id}", handlers.GetTaskByID)
	})

	log.Println("🚀 Server running on :8080")
	http.ListenAndServe(":8080", r)
}