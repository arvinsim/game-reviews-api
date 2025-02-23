package domain

import (
	"time"
)

type Review struct {
	ID          int       `json:"id"`
	GameID      int       `json:"gameID"`
	UserID      int       `json:"userID"`
	Rating      int       `json:"rating"`
	ReviewText  string    `json:"reviewText"`
	DateCreated time.Time `json:"dateCreated"`
}