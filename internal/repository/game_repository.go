package repository

import {
	"context"
}

type GameRepository interface {
    CreateGame(ctx context.Context, game *domain.Game) error
    GetGameByID(ctx context.Context, gameID int64) (*domain.Game, error)
}