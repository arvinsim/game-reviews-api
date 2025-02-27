package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/arvinsim/game-reviews-api/internal/domain"
	"github.com/arvinsim/game-reviews-api/internal/service"
)

type UserHandler interface {
	GetUsers(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) UserHandler {
	return &userHandler{service: service}
}

func (uh *userHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := uh.service.GetAllUsers(context.Background())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (uh *userHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
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
	_, err := uh.service.CreateUser(context.Background(), &newUser)
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
