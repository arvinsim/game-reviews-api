package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/arvinsim/game-reviews-api/internal/domain"
)

type GameHandler struct {}

func (gh *GameHandler) GetGames(w http.ResponseWriter, r *http.Request) {
	games := []domain.Game{
		{
			ID:          1,
			Title:       "Super Mario Bros",
			Description: "A classic game",
			Genre:       "Platformer",
			ReleaseDate: time.Date(1985, 9, 13, 0, 0, 0, 0, time.UTC),
			Developer:   "Nintendo",
		},
		{
			ID:          2,
			Title:       "The Legend of Zelda",
			Description: "Another classic game",
			Genre:       "Action-adventure",
			ReleaseDate: time.Date(1986, 2, 21, 0, 0, 0, 0, time.UTC),
			Developer:   "Nintendo",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(games); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}