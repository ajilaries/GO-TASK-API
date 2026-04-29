package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"go-task-api/internal/models"
	"go-task-api/internal/services"
)

// GET /tasks
func GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks := services.GetTasks()
	json.NewEncoder(w).Encode(tasks)
}

// POST /tasks
func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)

	created := services.CreateTask(task)
	json.NewEncoder(w).Encode(created)
}

// PUT /tasks/{id}
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/tasks/")
	id, _ := strconv.Atoi(idStr)

	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)

	updated, ok := services.UpdateTask(id, task)
	if !ok {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(updated)
}

// DELETE /tasks/{id}
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/tasks/")
	id, _ := strconv.Atoi(idStr)

	ok := services.DeleteTask(id)
	if !ok {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	w.Write([]byte("Deleted successfully"))
}