package repository

import (
	"context"
	"time"

	"com.ardafirdausr.cupid/internal/entity"
	"com.ardafirdausr.cupid/internal/entity/errs"
	"com.ardafirdausr.cupid/internal/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserMongoRepository struct {
	db *mongo.Database
}

func NewUserMongoRepository(db *mongo.Database) *UserMongoRepository {
	return &UserMongoRepository{db: db}
}

func (repo *UserMongoRepository) CreateUser(ctx context.Context, user *entity.User) error {
	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	res, err := repo.db.Collection("users").InsertOne(timeoutCtx, user)
	if err != nil {
		logger.Log.Err(err).Msg("failed to insert user")
		return errs.NewErrInternal("failed to insert user")
	}

	user.ID = res.InsertedID.(string)
	return err
}
