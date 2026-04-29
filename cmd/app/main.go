package main

import (
	"fmt"
	"net/http"

	"go-task-api/internal/handlers"
	"go-task-api/pkg/database"
)

func main() {
	database.ConnectDB()

	mux := http.NewServeMux()

	mux.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			handlers.CreateTask(w, r)
		} else if r.Method == http.MethodGet {
			handlers.GetTasks(w, r)
		}
	})

	fmt.Println("🚀 Server running on :8080")
	http.ListenAndServe(":8080", enableCORS(mux))
}
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}