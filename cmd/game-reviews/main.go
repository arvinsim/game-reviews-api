package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/arvinsim/game-reviews-api/internal/handlers"
)

func main() {
	mux := handlers.NewRouter()

	// Run the server
	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
