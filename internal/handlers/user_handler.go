package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/arvinsim/game-reviews-api/internal/domain"
)

type UserHandler struct {}

func (uh *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users := []domain.User{
		{1, "john123", "john123@gmail.com", "passwordhash123"},
		{2, "jane45", "jane456@gmail.com", "passwordhash456"},
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
