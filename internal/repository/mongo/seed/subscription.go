package seed

import (
	"context"

	"com.ardafirdausr.cupid/internal/entity"
	"com.ardafirdausr.cupid/internal/pkg/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var subscriptionPlans = []entity.Subscription{
	{
		ID:             entity.SubscriptionFree,
		Price:          0,
		DurationInDays: -1,
		Features: entity.SubscriptionFeature{
			MaxSwipe: 10,
			HasBadge: false,
		},
	},
	{
		ID:             entity.SubscriptionPremium,
		Price:          50000,
		DurationInDays: 30,
		Features: entity.SubscriptionFeature{
			MaxSwipe: -1,
			HasBadge: true,
		},
	}}

func SetupSubscriptionPlanCollection(ctx context.Context, database *mongo.Database) error {
	// seed the subscription plans
	upsertOption := options.Update().SetUpsert(true)
	for _, plan := range subscriptionPlans {
		_, err := database.Collection("subscriptions").UpdateOne(ctx, bson.M{"_id": plan.ID}, bson.M{"$set": plan}, upsertOption)
		if err != nil {
			logger.Log.Error().Err(err).Msg("failed to upsert subscription plan")
			return err
		}
	}

	_, err := database.Collection("user_subscriptions").
		Indexes().
		CreateMany(ctx, []mongo.IndexModel{
			{Keys: bson.M{"userID": 1}},
			{Keys: bson.M{"subscriptionID": 1}},
		})
	if err != nil {
		logger.Log.Error().Err(err).Msg("failed to upsert subscription plan")
		return err
	}

	return nil
}
