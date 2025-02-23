package domain

import (
	"time"
)

type Review struct {
	ID          int64       `json:"id"`
	GameID      int64       `json:"gameID"`
	UserID      int64       `json:"userID"`
	Rating      int8       `json:"rating"`
	ReviewText  string    `json:"reviewText"`
	DateCreated time.Time `json:"dateCreated"`
}