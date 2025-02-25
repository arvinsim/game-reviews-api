package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/arvinsim/game-reviews-api/internal/domain"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) error
	GetUserByID(ctx context.Context, userID int64) (*domain.User, error)
}

type userRepository struct {
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) CreateUser(ctx context.Context, user *domain.User) error {
	// implement DB or in-memory logic
	prettyJSON, err := json.MarshalIndent(user, "", "    ")
	if err != nil {
		fmt.Println("Failed to generate json", err)
		return err
	}
	fmt.Println(string(prettyJSON))

	return nil
}

func (r *userRepository) GetUserByID(ctx context.Context, userID int64) (*domain.User, error) {
	// implement DB or in-memory logic
	return &domain.User{
		ID:           userID,
		Username:     "john.doe",
		Email:        "john.doe@gmail.com",
		PasswordHash: "abc123"}, nil
}

func (r *userRepository) GetAllUsers(ctx context.Context) ([]*domain.User, error) {
	users := []*domain.User{
		{
			ID:           1,
			Username:     "john.doe",
			Email:        "john.doe@gmail.com",
			PasswordHash: "abc123",
		},
		{
			ID:           2,
			Username:     "jane.doe",
			Email:        "jane.doe@gmail.com",
			PasswordHash: "def456",
		},
	}
	return users, nil
}
