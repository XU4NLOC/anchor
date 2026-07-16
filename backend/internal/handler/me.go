package handler

import (
	"anchor-backend/internal/middleware"
	"anchor-backend/internal/user"
	"encoding/json"
	"log"
	"net/http"
)

type MeHandler struct {
	Repo *user.Repository
}

type meResponse struct {
	ID          string `json:"id"`
	Email       string `json:"email"`
	DisplayName string `json:"display_name"`
}

func (h *MeHandler) Me(w http.ResponseWriter, r *http.Request) {
	userID, ok := middleware.UserIDFromContext(r.Context())
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	foundUser, err := h.Repo.FindByID(r.Context(), userID)
	if err != nil {
		log.Printf("me: failed to find user: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	resp := meResponse{
		ID:          foundUser.ID,
		Email:       foundUser.Email,
		DisplayName: foundUser.DisplayName,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("me: failed to write response: %v", err)
	}
}
