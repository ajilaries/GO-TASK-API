package services

import (
	"errors"

	"go-task-api/internal/models"
	"go-task-api/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

// REGISTER
func Register(email, password string) error {

	if email == "" || password == "" {
		return errors.New("email and password required")
	}

	if len(password) < 6 {
		return errors.New("password must be at least 6 characters")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return err
	}

	user := models.User{
		Email:    email,
		Password: string(hashedPassword),
	}

	return repository.CreateUser(user)
}
//Login

func Login(email, password string)(string, error){
	if email==""|| password==""{
		return "",errors.New("Email and password already exists")
	}
	user, err:=repository.GetUserByEmail(email)
	if err!=nil{
		return "", errors.New("User not found")
	}
	err=bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	)
	if err!= nil{
		return "", errors.New("Invalid password")
	}
	//generate JWT
	token, err := GenerateToken(user.ID)
	if err!= nil{
		return "", err

	}
	return  token, nil
}