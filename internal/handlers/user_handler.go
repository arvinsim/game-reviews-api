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

	userResponses := convertToUserResponses(users)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(userResponses); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (uh *userHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUserRequest struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&newUserRequest); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	newUser := domain.User{
		Username:     newUserRequest.Username,
		Email:        newUserRequest.Email,
		PasswordHash: newUserRequest.Password,
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

func convertToUserResponses(users []*domain.User) []domain.UserResponse {
	userResponses := make([]domain.UserResponse, len(users))
	for i, user := range users {
		userResponses[i] = user.ConvertToUserResponse()
	}
	return userResponses
}
