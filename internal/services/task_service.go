package services

import (
	"go-task-api/internal/models"
	"go-task-api/internal/repository"
)

func GetTasks() []models.Task {
	return repository.GetAll()
}

func CreateTask(task models.Task) models.Task {
	return repository.Create(task)
}

func UpdateTask(id int, task models.Task) (models.Task, bool) {
	return repository.Update(id, task)
}

func DeleteTask(id int) bool {
	return repository.Delete(id)
}