package handlers

import (
	"encoding/json"
	"net/http"

	"go-task-api/internal/models"
	"go-task-api/internal/services"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)

	err := services.CreateTask(task)
	if err != nil {
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := services.GetTasks()
	if err != nil {
		http.Error(w, "Failed to fetch tasks", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(tasks)
}