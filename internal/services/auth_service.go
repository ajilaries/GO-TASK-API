package services

import (
	"errors"

	"go-task-api/internal/models"
	"go-task-api/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

// 🔐 use constant instead of magic number
const bcryptCost = 10

// REGISTER

func Register(email, password string)error{
	if email ==""|| password==""{
		return errors.New("Email and password can't be empty")
	}

	existingUser, err:= repository.GetUserByEmail(email)
	if err==nil && existingUser.ID!=""{
		return errors.New("Email already registered")
	}
	hashedPassword , err:= bcrypt.GenerateFromPassword([]byte(password),bcryptCost)
	if err!=nil{
		return err
	}
	user:=models.User{
		Email: email,
		Password: string(hashedPassword),

	}
	return repository.CreateUser(user)
}
// LOGIN
func Login(email, password string) (string, error) {

	if email == "" || password == "" {
		return "", errors.New("email and password required")
	}

	user, err := repository.GetUserByEmail(email)
	if err != nil {
		// 🔐 don't reveal if user exists
		return "", errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	)

	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// 🔥 generate JWT with user_id
	token, err := GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}