package handlers

import (
	"encoding/json"
	"net/http"

	"go-task-api/internal/services"
	"go-task-api/internal/utils"
)

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// REGISTER
func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req AuthRequest

	// ✅ Safe decode
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.Response{
			Error: "Invalid request body",
		})
		return
	}

	err := services.Register(req.Email, req.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.Response{
			Error: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(utils.Response{
		Message: "User registered successfully",
	})
}

// LOGIN
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req AuthRequest

	// ✅ Safe decode
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.Response{
			Error: "Invalid request body",
		})
		return
	}

	token, err := services.Login(req.Email, req.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(utils.Response{
			Error: err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(utils.Response{
		Token: token,
	})
}