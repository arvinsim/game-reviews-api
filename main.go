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

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("GET /users", getUsersHandler)
	mux.HandleFunc("GET /games", getGamesHandler)

	// Run the server
	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
