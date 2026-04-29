package main

import (
	"fmt"
	"log"
	"net/http"

	"go-task-api/internal/handlers"
	"go-task-api/pkg/database"

	"github.com/joho/godotenv"
)

func main() {

	//load .env file
	err:= godotenv.Load()
	if err !=nil{
		log.Println("No .env file found")
	}
	database.ConnectDB()

	mux := http.NewServeMux()

	mux.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			handlers.CreateTask(w, r)
		} else if r.Method == http.MethodGet {
			handlers.GetTasks(w, r)
		}
	})
	mux.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		handlers.DeleteTask(w, r)
	} else if r.Method == http.MethodPut {
		handlers.UpdateTask(w, r)
	}
})

	fmt.Println("🚀 Server running on :8080")
	http.ListenAndServe(":8080", enableCORS(mux))
}
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// 🔥 VERY IMPORTANT (this fixes your issue)
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}