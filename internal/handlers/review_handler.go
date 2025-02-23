package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/arvinsim/game-reviews-api/internal/domain"
)

type ReviewHandler struct {}

func (rh *ReviewHandler) GetReviews(w http.ResponseWriter, r *http.Request) {
	reviews := []domain.Review{
		{
			ID:          1,
			GameID:      1,
			UserID:      1,
			Rating:      5,
			ReviewText:  "Great game!",
			DateCreated: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			ID:          2,
			GameID:      1,
			UserID:      2,
			Rating:      4,
			ReviewText:  "Good game",
			DateCreated: time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC),
		},
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(reviews); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}