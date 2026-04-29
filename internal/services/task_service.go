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
func DeleteTask(id int)error{
	return repository.DeleteTask(id)
}
func UpdateTask(id int,title *string, completed *bool)error{
	return repository.UpdateTask(id,title,completed)
}