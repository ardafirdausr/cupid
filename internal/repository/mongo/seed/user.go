package seed

import (
	"context"
	"fmt"
	"time"

	"com.ardafirdausr.cupid/internal/entity"
	"com.ardafirdausr.cupid/internal/pkg/logger"
	"github.com/go-faker/faker/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func GenerateDummyUsers(ctx context.Context, database *mongo.Database, userCount int) error {
	batchSize := 200
	users := make([]interface{}, 0, batchSize)
	for i := 0; i < userCount; i++ {
		user := entity.User{}
		user.SetPassword("Rahasia123")
		user.ID = primitive.NewObjectID().Hex()
		user.Email = faker.Email()
		user.Bio = faker.Sentence()
		user.BirthDate, _ = time.Parse(time.DateOnly, faker.Date())

		user.Gender = entity.UserGenderMale
		user.Name = faker.FirstNameMale() + " " + faker.LastName()
		if i%2 == 0 {
			user.Gender = entity.UserGenderFemale
			user.Name = faker.FirstNameFemale() + " " + faker.LastName()
		}

		if i > 0 && i%batchSize == 0 || i == userCount-1 {
			logger.Log.Info().Msg(fmt.Sprintf("inserting %d users", len(users)))
			if _, err := database.Collection("users").InsertMany(ctx, users); err != nil {
				logger.Log.Error().Err(err).Msg("failed to insert users")
				return err
			}

			users = users[:0]
		} else {
			users = append(users, user)
		}

	}

	return nil
}
