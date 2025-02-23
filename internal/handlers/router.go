package handlers

import (
	"net/http"
)


func NewRouter() http.Handler  {
	mux := http.NewServeMux()

	// Maybe pass these as param?
	userHandler := &UserHandler{}
    gameHandler := &GameHandler{}
    reviewHandler := &ReviewHandler{}

	mux.HandleFunc("GET /users", userHandler.GetUsers)
	mux.HandleFunc("GET /games", gameHandler.GetGames)
	mux.HandleFunc("GET /reviews", reviewHandler.GetReviews)

	return mux
}