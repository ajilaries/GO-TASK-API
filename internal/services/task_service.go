package services

import (
	"go-task-api/internal/models"
	"go-task-api/internal/repository"
)

// CREATE
func CreateTask(task models.Task) error {
	return repository.CreateTask(task)
}

// GET ALL (user-specific)
func GetTasksByUser(userID string) ([]models.Task, error) {
	return repository.GetTasksByUser(userID)
}

// GET ONE
func GetTaskByID(id int, userID string) (*models.Task, error) {
	return repository.GetTaskByID(id, userID)
}

// DELETE (secure)
func DeleteTask(id int, userID string) error {
	return repository.DeleteTask(id, userID)
}

// UPDATE (secure)
func UpdateTask(id int, userID string, title *string, completed *bool) error {
	return repository.UpdateTask(id, userID, title, completed)
}