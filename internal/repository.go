//go:generate mockgen -source repository.go -package mock -destination ./mock/repository.go
package internal

import (
	"context"

	"com.ardafirdausr.cupid/internal/dto"
	"com.ardafirdausr.cupid/internal/entity"
)

type UserRepositorier interface {
	GetUserByID(ctx context.Context, id string) (*entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	CreateUser(ctx context.Context, user *entity.User) error
	UpdateUserByID(ctx context.Context, id string, user *entity.User) error
}

type MatchingRepositorier interface {
	GetMatchingRecommendations(ctx context.Context, filter dto.MatchingRecommendationsFilter) ([]entity.User, error)
	CreateMatching(ctx context.Context, matching *entity.Matching) error
}
