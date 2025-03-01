package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/arvinsim/game-reviews-api/internal/handlers"
	"github.com/arvinsim/game-reviews-api/internal/repository"
	"github.com/arvinsim/game-reviews-api/internal/service"
	"github.com/charmbracelet/lipgloss"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbPath := os.Getenv("SQLITE_DB_PATH")
	if dbPath == "" {
		log.Fatalf("SQLITE_DB_PATH is required")
	}

	db, err := sql.Open("sqlite3", dbPath)
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

	mux.HandleFunc("GET /users", corsMiddleware(userHandler.GetUsers))
	mux.HandleFunc("POST /users", userHandler.CreateUser)
	mux.HandleFunc("GET /games", gameHandler.GetGames)
	mux.HandleFunc("GET /reviews", reviewHandler.GetReviews)

	// Run the server
	runServer(8080, mux)
}

func runServer(port int, mux *http.ServeMux) {
	// Run the server
	message := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4")).
		Padding(0, 2)

	fmt.Printf(message.Render("Server is running on port: %d"), port)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}
