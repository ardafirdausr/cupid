package repository

import (
	"context"
	"time"

	"com.ardafirdausr.cupid/internal/entity"
	"com.ardafirdausr.cupid/internal/entity/errs"
	"com.ardafirdausr.cupid/internal/pkg/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserMongoRepository struct {
	db *mongo.Database
}

func NewUserMongoRepository(db *mongo.Database) *UserMongoRepository {
	return &UserMongoRepository{db: db}
}

func (repo *UserMongoRepository) GetUserByID(ctx context.Context, id string) (*entity.User, error) {
	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var user entity.User
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logger.Log.Err(err).Msg("failed to parse object id")
		return nil, errs.NewErrInternal("failed to parse object id")
	}

	err = repo.db.Collection("users").FindOne(timeoutCtx, primitive.M{"_id": objectID}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errs.NewErrNotFound("user not found")
		}

		logger.Log.Err(err).Msg("failed to get user by id")
		return nil, errs.NewErrInternal("failed to get user by id")
	}

	return &user, nil
}

func (repo *UserMongoRepository) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var user entity.User
	err := repo.db.Collection("users").FindOne(timeoutCtx, primitive.M{"email": email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errs.NewErrNotFound("user not found")
		}

		logger.Log.Err(err).Msg("failed to get user by email")
		return nil, errs.NewErrInternal("failed to get user by email")
	}

	return &user, nil

}

func (repo *UserMongoRepository) CreateUser(ctx context.Context, user *entity.User) error {
	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	user.ID = primitive.NewObjectID()
	_, err := repo.db.Collection("users").InsertOne(timeoutCtx, user)
	if err != nil {
		logger.Log.Err(err).Msg("failed to insert user")
		return errs.NewErrInternal("failed to insert user")
	}

	return err
}

func (repo *UserMongoRepository) UpdateUserByID(ctx context.Context, id string, user *entity.User) error {
	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logger.Log.Err(err).Msg("failed to parse object id")
		return errs.NewErrInternal("failed to parse object id")
	}

	_, err = repo.db.Collection("users").UpdateOne(timeoutCtx, primitive.M{"_id": objectID}, primitive.M{"$set": user})
	if err != nil {
		logger.Log.Err(err).Msg("failed to update user")
		return errs.NewErrInternal("failed to update user")
	}

	return nil
}