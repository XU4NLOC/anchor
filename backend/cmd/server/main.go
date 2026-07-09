package main

import (
	"log"
	"net/http"

	"anchor-backend/internal/router"
)

func main() {
	handler := router.New()

	log.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
