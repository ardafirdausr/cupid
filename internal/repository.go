//go:generate mockgen -source repository.go -package mock -destination ./mock/repository.go
package internal

import (
	"context"
	"time"

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
	GetUserMatchingCount(ctx context.Context, userID string, date time.Time) (uint64, error)
	GetMatchingByUser(ctx context.Context, user1ID, user2ID string) (*entity.Matching, error)
	CreateMatching(ctx context.Context, matching *entity.Matching) error
	UpdateMatchingByID(ctx context.Context, matchingID string, matching *entity.Matching) error
}

type SubscriptionRepositorier interface {
	GetAllSubscriptions(ctx context.Context) ([]entity.Subscription, error)
	GetSubscriptionByID(ctx context.Context, subscriptionID string) (*entity.Subscription, error)
	GetActiveUserSubscriptionByUserID(ctx context.Context, userID string) (*entity.UserSubscription, error)
	CreateUserSubscription(ctx context.Context, subscription *entity.UserSubscription) error
}
