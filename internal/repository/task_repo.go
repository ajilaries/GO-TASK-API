package repository

import (
	"go-task-api/internal/models"
	"go-task-api/pkg/database"
)

func CreateTask(task models.Task) error {
	query := "INSERT INTO tasks (title, completed) VALUES ($1, $2)"
	_, err := database.DB.Exec(query, task.Title, task.Completed)
	return err
}

func GetTasks() ([]models.Task, error) {
	rows, err := database.DB.Query("SELECT id, title, completed, created_at FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task

	for rows.Next() {
		var t models.Task
		err := rows.Scan(&t.ID, &t.Title, &t.Completed, &t.CreatedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	return tasks, nil
}