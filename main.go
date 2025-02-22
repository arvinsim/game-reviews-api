package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"passwordHash"`
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

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("GET /users", getUsersHandler)

	// Run the server
	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
