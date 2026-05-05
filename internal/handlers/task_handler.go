package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"go-task-api/internal/middleware"
	"go-task-api/internal/models"
	"go-task-api/internal/services"
	"go-task-api/internal/utils"
	"github.com/go-chi/chi/v5"
)


func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var task models.Task

	// ✅ Decode safely
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.Response{Error: "Invalid request body"})
		return
	}

	// 🔥 Get user from context
	userID := r.Context().Value(middleware.UserIDKey).(string)
	task.UserID = userID

	err := services.CreateTask(task)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.Response{Error: "Failed to create task"})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(utils.Response{Message: "Task created"})
}
func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userID := r.Context().Value(middleware.UserIDKey).(string)

	tasks, err := services.GetTasksByUser(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.Response{Error: "Failed to fetch tasks"})
		return
	}

	json.NewEncoder(w).Encode(utils.Response{Data: tasks})
}

func GetTaskByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userID := r.Context().Value(middleware.UserIDKey).(string)

	idStr := chi.URLParam(r, "id") // ✅ changed

	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.Response{Error: "Invalid task ID"})
		return
	}

	task, err := services.GetTaskByID(id, userID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(utils.Response{Error: "Task not found"})
		return
	}

	json.NewEncoder(w).Encode(utils.Response{Data: task})
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userID := r.Context().Value(middleware.UserIDKey).(string)

	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.Response{Error: "Invalid task ID"})
		return
	}

	err = services.DeleteTask(id, userID)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(utils.Response{Error: "Not allowed or task not found"})
		return
	}

	json.NewEncoder(w).Encode(utils.Response{Message: "Task deleted"})
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userID := r.Context().Value(middleware.UserIDKey).(string)

	idStr:=chi.URLParam(r,"id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.Response{Error: "Invalid task ID"})
		return
	}

	var req struct {
		Title     *string `json:"title"`
		Completed *bool   `json:"completed"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.Response{Error: "Invalid request body"})
		return
	}

	err = services.UpdateTask(id, userID, req.Title, req.Completed)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(utils.Response{Error: "Not allowed or task not found"})
		return
	}

	json.NewEncoder(w).Encode(utils.Response{Message: "Task updated"})
}