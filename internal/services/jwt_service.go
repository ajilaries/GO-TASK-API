package services

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey=[]byte("supersecretkey")

func GenerateToken(userID string)(string , error){
	token:= jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"user_id":userID,
		"exp":time.Now().Add(time.Hour*24).Unix(),//24 hr delay

	})
	return token.SignedString(secretKey)
}