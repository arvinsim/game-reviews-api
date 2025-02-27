package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/arvinsim/game-reviews-api/internal/domain"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/argon2"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	GetUserByID(ctx context.Context, userID int64) (*domain.User, error)
	GetAllUsers(ctx context.Context) ([]*domain.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	// implement DB or in-memory logic
	db, err := sql.Open("sqlite3", "../../data/game-reviews.db")
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	_, err = db.ExecContext(ctx, `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT NOT NULL UNIQUE,
			email TEXT NOT NULL UNIQUE,
			password_hash TEXT NOT NULL
		)
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to create users table: %v", err)
	}

	// Hash the password using Argon2id
	hashedPassword := argon2.IDKey([]byte(user.PasswordHash), []byte("somesalt"), 1, 64*1024, 4, 32)
	user.PasswordHash = fmt.Sprintf("%x", hashedPassword)

	result, err := db.ExecContext(ctx, `
		INSERT INTO users (username, email, password_hash)
		VALUES (?, ?, ?)
	`, user.Username, user.Email, user.PasswordHash)
	if err != nil {
		return nil, fmt.Errorf("failed to insert user: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to get last insert id: %v", err)
	}
	user.ID = id

	return user, nil
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
	rows, err := r.db.QueryContext(ctx, "SELECT id, username, email, password_hash FROM users")
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
