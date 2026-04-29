package services

import (
	"go-task-api/internal/models"
	"go-task-api/internal/repository"
)

func CreateTask(task models.Task)error{
	return repository.CreateTask(task)
}
func GetTasks() ([]models.Task , error){
	return repository.GetTasks()
}