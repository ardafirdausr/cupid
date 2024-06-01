//go:generate mockgen -source service.go -package mock -destination ./mock/service.go
package internal

import (
	"context"

	"com.ardafirdausr.cupid/internal/dto"
	"com.ardafirdausr.cupid/internal/entity"
)

type UserServicer interface {
	GetUserByID(ctx context.Context, id string) (*entity.User, error)
	UpdateUser(ctx context.Context, param dto.UpdateUserParam) (*entity.User, error)
}

type AuthServicer interface {
	Register(ctx context.Context, param dto.RegisterUserParam) (*entity.User, string, error)
	Login(ctx context.Context, param dto.LoginUserParam) (*entity.User, string, error)
}

type MatchingServicer interface {
	GetMatchingRecommendations(ctx context.Context, filter dto.MatchingRecommendationsFilter) ([]entity.User, error)
	MatchMaking(ctx context.Context, param dto.CreateMatchingParam) (*entity.Matching, error)
}

type SubscriptionServicer interface {
	GetAllSubscriptions(ctx context.Context) ([]entity.Subscription, error)
}
