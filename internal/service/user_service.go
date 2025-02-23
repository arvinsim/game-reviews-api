package service

import (
	"context"
	"github.com/arvinsim/game-reviews-api/internal/domain"
	"github.com/arvinsim/game-reviews-api/internal/repository"
)

type UserService interface {
	CreateUser(ctx context.Context, name string) (*domain.User, error)
    GetUser(ctx context.Context, userID int64) (*domain.User, error)
    GetAllUsers(ctx context.Context, userID int64) (*domain.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (*us UserService) GetUser(ctx context.Context, userID int64) (*domain.User, error) {
	return us.userRepo.GetUser(ctx, userID)
}