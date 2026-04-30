package repository

import (
	"go-task-api/internal/models"
	"go-task-api/pkg/database"
)

func CreateUser(user models.User) error {
	query := "INSERT INTO users (email, password) VALUES ($1, $2)"
	_, err := database.DB.Exec(query, user.Email, user.Password)
	return err
}

func GetUserByEmail(email string) (models.User, error) {
	var user models.User
	query := "SELECT id, email, password FROM users WHERE email=$1"

	err := database.DB.QueryRow(query, email).
		Scan(&user.ID, &user.Email, &user.Password)

	return user, err
}