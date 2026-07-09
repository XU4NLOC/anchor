package main

import (
	"log"
	"net/http"

	"anchor-backend/internal/config"
	"anchor-backend/internal/router"
)

func main() {
	cfg := config.Load()
	handler := router.New()

	log.Printf("Server is running on port: %v", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, handler); err != nil {
		log.Fatal(err)
	}
}
