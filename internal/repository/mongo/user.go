package mongo

import (
	"context"
	"time"

	"com.ardafirdausr.cupid/internal/entity"
	"com.ardafirdausr.cupid/internal/entity/errs"
	"com.ardafirdausr.cupid/internal/pkg/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	db *mongo.Database
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{db: db}
}

func (repo *UserRepository) GetUserByID(ctx context.Context, id string) (*entity.User, error) {
	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var user entity.User
	if err := repo.db.Collection("users").FindOne(timeoutCtx, bson.M{"_id": id}).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errs.NewErrNotFound("user not found")
		}

		logger.Log.Err(err).Msg("failed to get user by id")
		return nil, errs.NewErrInternal("failed to get user by id")
	}

	return &user, nil
}

func (repo *UserRepository) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var user entity.User
	err := repo.db.Collection("users").FindOne(timeoutCtx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errs.NewErrNotFound("user not found")
		}

		logger.Log.Err(err).Msg("failed to get user by email")
		return nil, errs.NewErrInternal("failed to get user by email")
	}

	return &user, nil

}

func (repo *UserRepository) CreateUser(ctx context.Context, user *entity.User) error {
	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	user.ID = primitive.NewObjectID().Hex()
	_, err := repo.db.Collection("users").InsertOne(timeoutCtx, user)
	if err != nil {
		logger.Log.Err(err).Msg("failed to create user")
		return errs.NewErrInternal("failed to create user")
	}

	return nil
}

func (repo *UserRepository) UpdateUserByID(ctx context.Context, id string, user *entity.User) error {
	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if _, err := repo.db.Collection("users").UpdateOne(timeoutCtx, bson.M{"_id": user.Bio}, bson.M{"$set": user}); err != nil {
		logger.Log.Err(err).Msg("failed to update user")
		return errs.NewErrInternal("failed to update user")
	}

	return nil
}
