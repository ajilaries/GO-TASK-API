package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"go-task-api/internal/models"
	"go-task-api/internal/services"
	"github.com/gorilla/mux"
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

func GetTaskByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]

	// TODO: replace with DB/service call
	task := map[string]string{
		"id":    id,
		"title": "Sample Task",
	}

	json.NewEncoder(w).Encode(task)
}
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/tasks/")
	id, _ := strconv.Atoi(idStr)

	err := services.DeleteTask(id)
	if err != nil {
		http.Error(w, "Failed to delete task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/tasks/")
	id, _ := strconv.Atoi(idStr)

	var req struct {
		Title     *string `json:"title"`
		Completed *bool   `json:"completed"`
	}

	json.NewDecoder(r.Body).Decode(&req)

	err := services.UpdateTask(id, req.Title, req.Completed)
	if err != nil {
		http.Error(w, "Failed to update task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}