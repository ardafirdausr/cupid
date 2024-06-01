package seed

import (
	"context"

	"com.ardafirdausr.cupid/internal/pkg/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupMatchingCollection(ctx context.Context, database *mongo.Database) error {
	_, err := database.Collection("users").
		Indexes().
		CreateMany(ctx, []mongo.IndexModel{
			{Keys: bson.D{{Key: "user1ID", Value: 1}, {Key: "status", Value: 1}}},
			{Keys: bson.D{{Key: "user2ID", Value: 1}, {Key: "status", Value: 1}}},
		})
	if err != nil {
		logger.Log.Error().Err(err).Msg("failed to create indexes for matching collection")
		return err
	}

	return nil
}
