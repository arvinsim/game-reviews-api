package repository

import (
	"context"

	"github.com/arvinsim/game-reviews-api/internal/domain"
)

type GameRepository interface {
	CreateGame(ctx context.Context, game *domain.Game) error
	GetGameByID(ctx context.Context, gameID int64) (*domain.Game, error)
}
