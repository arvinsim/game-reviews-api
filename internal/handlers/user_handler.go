package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/arvinsim/game-reviews-api/internal/domain"
	"github.com/arvinsim/game-reviews-api/internal/repository"
	"github.com/arvinsim/game-reviews-api/internal/service"
)

type UserHandler struct{}

func (uh *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users := []domain.User{
		{ID: 1, Username: "john.doe", Email: "john123@gmail.com", PasswordHash: "passwordhash123"},
		{ID: 2, Username: "jane.doe", Email: "jane456@gmail.com", PasswordHash: "passwordhash456"},
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (uh *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newUser := domain.User{
		Username:     r.FormValue("username"),
		Email:        r.FormValue("email"),
		PasswordHash: r.FormValue("password_hash"),
	}

	// Here you would typically add code to save the new user to a database
	// For now, we'll just return the user as a confirmation
	userRepo := repository.NewUserRepository() // Assuming you have a NewUserRepository function
	userService := service.NewUserService(userRepo)
	_, err := userService.CreateUser(context.Background(), &newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Create User failed 1, error: ", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(newUser); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Create User failed 2, error: ", err)
	}
}
