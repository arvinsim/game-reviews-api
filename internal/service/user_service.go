package service

import (
	"context"

	"github.com/arvinsim/game-reviews-api/internal/domain"
	"github.com/arvinsim/game-reviews-api/internal/repository"
)

type UserService interface {
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	GetUser(ctx context.Context, userID int64) (*domain.User, error)
	GetAllUsers(ctx context.Context) ([]*domain.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (us *userService) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	_, err := us.userRepo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *userService) GetUser(ctx context.Context, userID int64) (*domain.User, error) {
	return us.userRepo.GetUserByID(ctx, userID)
}

func (us *userService) GetAllUsers(ctx context.Context) ([]*domain.User, error) {
	return us.userRepo.GetAllUsers(ctx)
}
