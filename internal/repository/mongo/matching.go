package mongo

import (
	"context"
	"time"

	"com.ardafirdausr.cupid/internal/dto"
	"com.ardafirdausr.cupid/internal/entity"
	"com.ardafirdausr.cupid/internal/entity/errs"
	"com.ardafirdausr.cupid/internal/pkg/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MatchingMongoRepositry struct {
	db *mongo.Database
}

func NewMatchingMongoRepository(db *mongo.Database) *MatchingMongoRepositry {
	return &MatchingMongoRepositry{db: db}
}

func (repo *MatchingMongoRepositry) GetMatchingRecommendations(ctx context.Context, filter dto.MatchingRecommendationsFilter) ([]entity.User, error) {
	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	pipeline := mongo.Pipeline{
		bson.D{
			{Key: "$match", Value: bson.D{
				{Key: "_id", Value: bson.D{
					{Key: "$ne", Value: filter.UserID},
				}},
			}},
		},
		bson.D{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "matchings"},
				{Key: "localField", Value: "_id"},
				{Key: "foreignField", Value: "user1ID"},
				{Key: "as", Value: "user1Matchings"},
			}},
		},
		bson.D{
			{Key: "$match", Value: bson.D{
				{Key: "user1Matchings.user1ID", Value: bson.D{
					{Key: "$exists", Value: false},
				}},
			}},
		},
		bson.D{{Key: "$limit", Value: 10}},
	}

	cursor, err := repo.db.Collection("users").Aggregate(timeoutCtx, pipeline)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errs.NewErrNotFound("user not found")
		}

		logger.Log.Err(err).Msg("failed to get user by id")
		return nil, errs.NewErrInternal("failed to get user by id")
	}

	defer cursor.Close(timeoutCtx)

	// Iterate over results
	users := make([]entity.User, 0)
	for cursor.Next(context.TODO()) {
		var user entity.User
		if err := cursor.Decode(&user); err != nil {
			logger.Log.Err(err).Msg("failed to decode user")
			return nil, errs.NewErrInternal("failed to decode user")
		}

		users = append(users, user)
	}

	return users, nil
}

func (repo *MatchingMongoRepositry) CreateMatching(ctx context.Context, matching *entity.Matching) error {
	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	matching.ID = primitive.NewObjectID().Hex()
	if _, err := repo.db.Collection("matchings").InsertOne(timeoutCtx, matching); err != nil {
		logger.Log.Err(err).Msg("failed to create matching")
		return errs.NewErrInternal("failed to create matching")
	}

	return nil
}
