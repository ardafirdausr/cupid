package seed

import (
	"context"

	"com.ardafirdausr.cupid/internal/entity"
	"com.ardafirdausr.cupid/internal/pkg/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var subscriptionPlans = []entity.SubscriptionPlan{
	{
		ID:             entity.SubscriptionTypeFree,
		Price:          0,
		DurationInDays: -1,
		Features: entity.SubscriptionFeature{
			entity.SubscriptionFeatureMaxSwipe: 10,
			entity.SubscriptionFeatureHasBadge: false,
		},
	},
	{
		ID:             entity.SubscriptionTypePremium,
		Price:          50000,
		DurationInDays: 30,
		Features: entity.SubscriptionFeature{
			entity.SubscriptionFeatureMaxSwipe: -1,
			entity.SubscriptionFeatureHasBadge: true,
		},
	}}

func SetupSubscriptionPlanCollection(ctx context.Context, database *mongo.Database) error {
	// seed the subscription plans
	collection := database.Collection("subscription_plans")

	// upsert the plans
	upsertOption := options.Update().SetUpsert(true)
	for _, plan := range subscriptionPlans {
		_, err := collection.UpdateOne(ctx, bson.M{"_id": plan.ID}, bson.M{"$set": plan}, upsertOption)
		if err != nil {
			logger.Log.Error().Err(err).Msg("failed to upsert subscription plan")
			return err
		}
	}

	return nil
}
