package mongo

import (
	"context"
	"time"

	"com.ardafirdausr.cupid/internal/entity/errs"
	"com.ardafirdausr.cupid/internal/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDatabase(config Config) (*mongo.Database, func(), error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	options := options.Client().ApplyURI(config.URI).
		SetMinPoolSize(config.MinPool).
		SetMaxPoolSize(config.MaxPool).
		SetConnectTimeout(config.MaxIdleTimePool)
	client, err := mongo.Connect(ctx, options)
	if err != nil {
		logger.Log.Err(err).Msg("failed to connect to mongo")
		return nil, func() {}, errs.NewErrInternal("failed to connect to mongo")
	}

	if err = client.Ping(ctx, nil); err != nil {
		logger.Log.Err(err).Msg("failed to ping mongo")
		return nil, func() {}, errs.NewErrInternal("failed to ping mongo")
	}

	return client.Database(config.DB), func() {
		closeCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := client.Disconnect(closeCtx); err != nil {
			logger.Log.Err(err).Msg("failed to disconnect from mongo")
		}
	}, nil
}
