package repository

import {
	"context"
}

type ReviewRepository interface {
    CreateReview(ctx context.Context, review *domain.Review) error
    GetReviewByID(ctx context.Context, reviewID int64) (*domain.Review, error)
}