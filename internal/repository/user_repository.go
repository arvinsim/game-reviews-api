package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/arvinsim/game-reviews-api/internal/domain"

	_ "github.com/mattn/go-sqlite3"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) error
	GetUserByID(ctx context.Context, userID int64) (*domain.User, error)
	GetAllUsers(ctx context.Context) ([]*domain.User, error)
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
	db, err := sql.Open("sqlite3", "../../data/game-reviews.db")
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}
	defer db.Close()

	rows, err := db.QueryContext(ctx, "SELECT id, username, email, password_hash FROM users")
	if err != nil {
		return nil, fmt.Errorf("failed to query users: %w", err)
	}
	defer rows.Close()

	var users []*domain.User
	for rows.Next() {
		var user domain.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash); err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return users, nil
}
