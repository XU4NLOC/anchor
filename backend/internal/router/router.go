package router

import (
	"anchor-backend/internal/handler"
	"anchor-backend/internal/user"
	"net/http"
)

func New(repo *user.Repository, jwtSecret string) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", handler.Health)

	registerHandler := &handler.RegisterHandler{Repo: repo}
	mux.HandleFunc("POST /register", registerHandler.Register)

	loginHandler := &handler.LoginHandler{Repo: repo, JWTSecret: jwtSecret}
	mux.HandleFunc("POST /login", loginHandler.Login)

	return mux
}
