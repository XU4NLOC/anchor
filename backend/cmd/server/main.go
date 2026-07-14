package main

import (
	"anchor-backend/internal/config"
	"anchor-backend/internal/database"
	"anchor-backend/internal/router"
	"anchor-backend/internal/user"
	"context"
	"log"
	"net/http"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	ctx := context.Background()

	pool, err := database.New(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer pool.Close()

	log.Printf("Connected to database")

	repo := user.NewRepository(pool)

	handler := router.New(repo, cfg.JWTSecret)

	log.Printf("Server is running on port: %v", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, handler); err != nil {
		log.Fatal(err)
	}
}
