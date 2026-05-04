package repository

import (
	"fmt"
	"go-task-api/internal/models"
	"go-task-api/pkg/database"
)

func CreateTask(task models.Task) error {
	query := "INSERT INTO tasks (title, completed, user_id) VALUES ($1, $2, $3)"
	_, err := database.DB.Exec(query, task.Title, task.Completed, task.UserID)
	return err
}
func GetTasksByUser(userID string) ([]models.Task, error) {

	rows, err := database.DB.Query(
		"SELECT id, title, completed, created_at FROM tasks WHERE user_id=$1",
		userID,
	)
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
func GetTaskByID(id int, userID string) (*models.Task, error) {

	var t models.Task

	err := database.DB.QueryRow(
		"SELECT id, title, completed, created_at FROM tasks WHERE id=$1 AND user_id=$2",
		id, userID,
	).Scan(&t.ID, &t.Title, &t.Completed, &t.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &t, nil
}
func DeleteTask(id int, userID string) error {

	result, err := database.DB.Exec(
		"DELETE FROM tasks WHERE id=$1 AND user_id=$2",
		id, userID,
	)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("not found or not allowed")
	}

	return nil
}
func UpdateTask(id int, userID string, title *string, completed *bool) error {

	query := "UPDATE tasks SET "
	params := []interface{}{}
	i := 1

	if title != nil {
		query += "title=$" + fmt.Sprint(i)
		params = append(params, *title)
		i++
	}

	if completed != nil {
		if len(params) > 0 {
			query += ", "
		}
		query += "completed=$" + fmt.Sprint(i)
		params = append(params, *completed)
		i++
	}

	// 🔥 IMPORTANT: secure condition
	query += " WHERE id=$" + fmt.Sprint(i) + " AND user_id=$" + fmt.Sprint(i+1)
	params = append(params, id, userID)

	result, err := database.DB.Exec(query, params...)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("not found or not allowed")
	}

	return nil
}