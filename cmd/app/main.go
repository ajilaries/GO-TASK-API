package main

import (
	"fmt"
	"net/http"

	"go-task-api/internal/handlers"
)

func main() {
	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handlers.GetTasks(w, r)
		} else if r.Method == "POST" {
			handlers.CreateTask(w, r)
		}
	})

	http.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "PUT" {
			handlers.UpdateTask(w, r)
		} else if r.Method == "DELETE" {
			handlers.DeleteTask(w, r)
		}
	})

	fmt.Println("Server running on :8080 🚀")
	http.ListenAndServe(":8080", nil)
}