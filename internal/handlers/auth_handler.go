package handlers

import (
	"encoding/json"
	"net/http"

	"go-task-api/internal/services"
)

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// REGISTER
func Register(w http.ResponseWriter, r *http.Request) {
	var req AuthRequest
	json.NewDecoder(r.Body).Decode(&req)

	err := services.Register(req.Email, req.Password)
	if err != nil {
		http.Error(w, "Failed to register", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("User registered successfully"))
}

// LOGIN
func Login(w http.ResponseWriter, r *http.Request) {
	var req AuthRequest
	json.NewDecoder(r.Body).Decode(&req)

	token, err := services.Login(req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}