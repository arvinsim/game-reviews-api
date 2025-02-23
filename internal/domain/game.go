package domain

import (
	"time"
)

type Game struct {
	ID          int64       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Genre       string    `json:"genre"`
	ReleaseDate time.Time `json:"releaseDate"`
	Developer   string    `json:"developer"`
}