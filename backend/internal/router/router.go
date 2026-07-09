package router

import (
	"net/http"

	"anchor-backend/internal/handler"
)

func New() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", handler.Health)
	return mux
}
