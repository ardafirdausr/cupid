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

type MatchingRepository struct {
	db *mongo.Database
}

func NewMatchingRepository(db *mongo.Database) *MatchingRepository {
	return &MatchingRepository{db: db}
}

func (repo *MatchingRepository) GetUserMatchingCount(ctx context.Context, userID string, date time.Time) (uint64, error) {
	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	today := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	tommorow := today.AddDate(0, 0, 1)
	pipeline := mongo.Pipeline{
		bson.D{
			{Key: "$match", Value: bson.D{
				{Key: "$and", Value: bson.A{
					bson.D{{Key: "user1ID", Value: userID}},
					bson.D{{Key: "user1SwapAt", Value: bson.D{
						{Key: "$gte", Value: today},
						{Key: "$lt", Value: tommorow},
					}}},
				}},
			}},
		},
		bson.D{
			{Key: "$count", Value: "count"},
		},
	}

	cursor, err := repo.db.Collection("matchings").Aggregate(timeoutCtx, pipeline)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return 0, nil
		}

		logger.Log.Err(err).Msg("failed to get user accepted count")
	}

	defer cursor.Close(timeoutCtx)

	var result struct {
		Count uint64 `bson:"count"`
	}

	if cursor.Next(context.Background()) {
		if err := cursor.Decode(&result); err != nil {
			logger.Log.Err(err).Msg("failed to decode user accepted count")
			return 0, errs.NewErrInternal("failed to decode user accepted count")
		}
	}

	return result.Count, nil
}

func (repo *MatchingRepository) GetMatchingRecommendations(ctx context.Context, filter dto.MatchingRecommendationsFilter) ([]entity.User, error) {
	timeoutCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	pipeline := mongo.Pipeline{
		bson.D{
			{Key: "$match", Value: bson.D{
				{Key: "$and", Value: bson.A{
					bson.D{{Key: "_id", Value: bson.D{{Key: "$ne", Value: filter.UserID}}}},
					bson.D{{Key: "gender", Value: bson.D{{Key: "$eq", Value: filter.Gender}}}},
				}},
			}},
		},
		bson.D{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "matchings"},
				{Key: "as", Value: "matching"},
				{Key: "let", Value: bson.D{{Key: "userID", Value: "$_id"}}},
				{Key: "pipeline", Value: bson.A{
					bson.D{{Key: "$match", Value: bson.D{
						{Key: "$expr", Value: bson.D{
							{Key: "$or", Value: bson.A{
								bson.D{{Key: "$and", Value: bson.A{ // user's matchings sent
									bson.D{{Key: "$eq", Value: bson.A{"$user1ID", filter.UserID}}},
									bson.D{{Key: "$eq", Value: bson.A{"$user2ID", "$$userID"}}},
								}}},
								bson.D{{Key: "$and", Value: bson.A{ // user's matchings received but not matched or rejected yet
									bson.D{{Key: "$eq", Value: bson.A{"$user1ID", "$$userID"}}},
									bson.D{{Key: "$eq", Value: bson.A{"$user2ID", filter.UserID}}},
									bson.D{{Key: "$in", Value: bson.A{"$status", []entity.MatchingStatus{entity.MatchingStatusMatched, entity.MatchingStatusRejected}}}},
								}}},
							}},
						}},
					}}},
				}},
			}},
		},
		bson.D{ // ignore users that have 2 conditions above
			{Key: "$match", Value: bson.D{
				{Key: "matching", Value: bson.D{
					{Key: "$eq", Value: bson.A{}},
				}},
			}},
		},
		bson.D{{Key: "$limit", Value: filter.Limit}},
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

func (repo *MatchingRepository) GetMatchingByUser(ctx context.Context, user1ID, user2ID string) (*entity.Matching, error) {
	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var matching entity.Matching
	if err := repo.db.Collection("matchings").
		FindOne(timeoutCtx, bson.M{
			"$or": bson.A{
				bson.M{"user1ID": user1ID, "user2ID": user2ID},
				bson.M{"user1ID": user2ID, "user2ID": user1ID},
			},
		}).
		Decode(&matching); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errs.NewErrNotFound("matching not found")
		}

		logger.Log.Err(err).Msg("failed to get matching by user")
		return nil, errs.NewErrInternal("failed to get matching by user")
	}

	return &matching, nil
}

func (repo *MatchingRepository) CreateMatching(ctx context.Context, matching *entity.Matching) error {
	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	matching.ID = primitive.NewObjectID().Hex()
	if _, err := repo.db.Collection("matchings").InsertOne(timeoutCtx, matching); err != nil {
		logger.Log.Err(err).Msg("failed to create matching")
		return errs.NewErrInternal("failed to create matching")
	}

	return nil
}

func (repo *MatchingRepository) UpdateMatchingByID(ctx context.Context, matchingID string, matching *entity.Matching) error {
	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := repo.db.Collection("matchings").UpdateOne(timeoutCtx, bson.M{"_id": matchingID}, bson.M{"$set": matching})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errs.NewErrNotFound("matching not found")
		}

		logger.Log.Err(err).Msg("failed to update matching")
		return errs.NewErrInternal("failed to update matching")
	}

	return nil
}
