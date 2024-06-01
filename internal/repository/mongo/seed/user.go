package seed

import (
	"context"

	"com.ardafirdausr.cupid/internal/pkg/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SetupUserCollection(ctx context.Context, database *mongo.Database) error {
	// create unique index to email
	_, err := database.Collection("users").
		Indexes().
		CreateOne(ctx, mongo.IndexModel{
			Keys:    bson.M{"email": 1},
			Options: options.Index().SetUnique(true),
		})
	if err != nil {
		logger.Log.Error().Err(err).Msg("failed to create unique index to email")
		return err
	}

	return nil
}
