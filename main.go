package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type User struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"passwordHash"`
}

type Game struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Genre       string    `json:"genre"`
	ReleaseDate time.Time `json:"releaseDate"`
	Developer   string    `json:"developer"`
}

type Review struct {
	ID          int       `json:"id"`
	GameID      int       `json:"gameID"`
	UserID      int       `json:"userID"`
	Rating      int       `json:"rating"`
	ReviewText  string    `json:"reviewText"`
	DateCreated time.Time `json:"dateCreated"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is the root web page")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	users := []User{
		{1, "john123", "john123@gmail.com", "passwordhash123"},
		{2, "jane45", "jane456@gmail.com", "passwordhash456"},
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getGamesHandler(w http.ResponseWriter, r *http.Request) {
	games := []Game{
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

func getReviewsHandler(w http.ResponseWriter, r *http.Request) {
	reviews := []Review{
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

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("GET /users", getUsersHandler)
	mux.HandleFunc("GET /games", getGamesHandler)
	mux.HandleFunc("GET /reviews", getReviewsHandler)

	// Run the server
	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
