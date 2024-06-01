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

func (repo *SubscriptionRepository) GetSubscriptionByID(ctx context.Context, subscriptionID string) (*entity.Subscription, error) {
	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var subscription entity.Subscription
	var filter = bson.M{"_id": subscriptionID}
	if err := repo.DB.Collection("subscriptions").FindOne(timeoutCtx, filter).Decode(&subscription); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errs.NewErrNotFound("subscription not found")
		}

		logger.Log.Err(err).Msg("failed to get subscription by id")
		return nil, errs.NewErrInternal("failed to get subscription by id")
	}

	return &subscription, nil
}

func (repo *SubscriptionRepository) GetActiveUserSubscriptionByUserID(ctx context.Context, userID string) (*entity.UserSubscription, error) {
	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var userSubscription entity.UserSubscription
	var filter = bson.M{"userID": userID, "expiredAt": bson.M{"$gt": time.Now()}}
	if err := repo.DB.Collection("user_subscriptions").FindOne(timeoutCtx, filter).Decode(&userSubscription); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errs.NewErrNotFound("subscription not found")
		}

		logger.Log.Err(err).Msg("failed to get subscription by id")
		return nil, errs.NewErrInternal("failed to get subscription by id")
	}

	return &userSubscription, nil
}

func (repo *SubscriptionRepository) CreateUserSubscription(ctx context.Context, subscription *entity.UserSubscription) error {
	return nil
}
