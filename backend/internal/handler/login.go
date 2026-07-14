package handler

import "anchor-backend/internal/user"

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
