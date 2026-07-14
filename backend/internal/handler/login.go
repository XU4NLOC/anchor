package handler

import (
	"anchor-backend/internal/password"
	"anchor-backend/internal/token"
	"anchor-backend/internal/user"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
)

type LoginHandler struct {
	Repo      *user.Repository
	JWTSecret string
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginResponse struct {
	Token string `json:"token"`
}

func (h *LoginHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req loginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Email == "" || req.Password == "" {
		http.Error(w, "Email and password are required", http.StatusBadRequest)
		return
	}

	foundUser, err := h.Repo.FindByEmail(r.Context(), req.Email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
			return
		}

		log.Printf("login: failed to find user: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if !password.Verify(req.Password, foundUser.PasswordHash) {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	signed, err := token.Generate(foundUser.ID, h.JWTSecret)
	if err != nil {
		log.Printf("login: failed to generate tolen: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(loginResponse{Token: signed}); err != nil {
		log.Printf("login: failed to write response: %v", err)
	}
}
