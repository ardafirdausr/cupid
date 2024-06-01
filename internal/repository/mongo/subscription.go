package mongo

import (
	"context"
	"time"

	"com.ardafirdausr.cupid/internal/entity"
	"com.ardafirdausr.cupid/internal/entity/errs"
	"com.ardafirdausr.cupid/internal/pkg/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type SubscriptionRepository struct {
	DB *mongo.Database
}

func NewSubscriptionRepository(db *mongo.Database) *SubscriptionRepository {
	return &SubscriptionRepository{DB: db}
}

func (repo *SubscriptionRepository) GetAllSubscriptions(ctx context.Context) ([]entity.Subscription, error) {
	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	subscriptions := make([]entity.Subscription, 0)

	cursor, err := repo.DB.Collection("subscriptions").Find(timeoutCtx, bson.M{})
	if err != nil {
		logger.Log.Err(err).Msg("failed to query subscriptions")
		return subscriptions, errs.NewErrInternal("failed to query subscriptions")
	}

	defer cursor.Close(timeoutCtx)

	for cursor.Next(timeoutCtx) {
		var subscription entity.Subscription
		if err := cursor.Decode(&subscription); err != nil {
			logger.Log.Err(err).Msg("failed to decode subscription")
			return subscriptions, errs.NewErrInternal("failed to decode subscription")
		}

		subscriptions = append(subscriptions, subscription)
	}

	return subscriptions, nil
}
