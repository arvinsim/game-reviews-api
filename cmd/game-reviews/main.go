package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/arvinsim/game-reviews-api/internal/handlers"
	"github.com/arvinsim/game-reviews-api/internal/repository"
	"github.com/arvinsim/game-reviews-api/internal/service"
)

func main() {
	db, err := sql.Open("sqlite3", "../../data/game-reviews.db")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	mux := http.NewServeMux()
	gameHandler := &handlers.GameHandler{}
	reviewHandler := &handlers.ReviewHandler{}

	mux.HandleFunc("GET /users", userHandler.GetUsers)
	mux.HandleFunc("POST /users", userHandler.CreateUser)
	mux.HandleFunc("GET /games", gameHandler.GetGames)
	mux.HandleFunc("GET /reviews", reviewHandler.GetReviews)

	// Run the server
	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
