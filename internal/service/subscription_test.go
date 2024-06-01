package service

import (
	"context"
	"testing"

	"com.ardafirdausr.cupid/internal/dto"
	"com.ardafirdausr.cupid/internal/entity"
	"com.ardafirdausr.cupid/internal/entity/errs"
	"com.ardafirdausr.cupid/internal/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetAllSubscriptions(t *testing.T) {
	t.Run("failed when getting all subscriptions", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		subscriptionRepo := mock.NewMockSubscriptionRepositorier(ctrl)
		subscriptionRepo.EXPECT().GetAllSubscriptions(gomock.Any()).Return(nil, errs.ErrInternal).Times(1)

		subscriptionService := NewSubscriptionService(subscriptionRepo)
		subscriptions, err := subscriptionService.GetAllSubscriptions(context.Background())
		assert.Nil(t, subscriptions)
		assert.NotNil(t, err)
		assert.True(t, errs.IsEqualType(err, errs.ErrInternal))
	})

	t.Run("success when getting all subscriptions", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		subscriptions := []entity.Subscription{
			{
				ID: "1",
			},
			{
				ID: "2",
			},
		}

		subscriptionRepo := mock.NewMockSubscriptionRepositorier(ctrl)
		subscriptionRepo.EXPECT().GetAllSubscriptions(gomock.Any()).Return(subscriptions, nil).Times(1)

		subscriptionService := NewSubscriptionService(subscriptionRepo)
		result, err := subscriptionService.GetAllSubscriptions(context.Background())
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, len(result), 2)
	})
}

func TestGetActiveUserSubscription(t *testing.T) {
	t.Run("failed when getting user from context", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		subscriptionRepo := mock.NewMockSubscriptionRepositorier(ctrl)
		subscriptionService := NewSubscriptionService(subscriptionRepo)
		userSubscription, err := subscriptionService.GetActiveUserSubscription(context.Background())
		assert.Nil(t, userSubscription)
		assert.NotNil(t, err)
		assert.True(t, errs.IsEqualType(err, errs.ErrInternal))
	})

	t.Run("failed when getting active user subscription by user id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		user := entity.User{
			ID: "1",
		}

		subscriptionRepo := mock.NewMockSubscriptionRepositorier(ctrl)
		subscriptionRepo.EXPECT().GetActiveUserSubscriptionByUserID(gomock.Any(), user.ID).Return(nil, errs.ErrInternal).Times(1)

		subscriptionService := NewSubscriptionService(subscriptionRepo)
		userSubscription, err := subscriptionService.GetActiveUserSubscription(context.WithValue(context.Background(), entity.UserContextKey, &user))
		assert.Nil(t, userSubscription)
		assert.NotNil(t, err)
		assert.True(t, errs.IsEqualType(err, errs.ErrInternal))
	})

	t.Run("success when getting active user subscription by user id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		user := entity.User{
			ID: "1",
		}

		userSubscription := entity.UserSubscription{
			UserID: "1",
		}

		subscriptionRepo := mock.NewMockSubscriptionRepositorier(ctrl)
		subscriptionRepo.EXPECT().GetActiveUserSubscriptionByUserID(gomock.Any(), user.ID).Return(&userSubscription, nil).Times(1)

		subscriptionService := NewSubscriptionService(subscriptionRepo)
		result, err := subscriptionService.GetActiveUserSubscription(context.WithValue(context.Background(), entity.UserContextKey, &user))
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, result.UserID, user.ID)
	})
}

func TestCreateUserSubscription(t *testing.T) {
	t.Run("failed when getting user from context", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		subscriptionRepo := mock.NewMockSubscriptionRepositorier(ctrl)
		subscriptionService := NewSubscriptionService(subscriptionRepo)
		userSubscription, err := subscriptionService.CreateUserSubscription(context.Background(), dto.CreateUserSubscriptionParam{})
		assert.Nil(t, userSubscription)
		assert.NotNil(t, err)
		assert.True(t, errs.IsEqualType(err, errs.ErrInternal))
	})

	t.Run("failed when subscription id is invalid", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		user := entity.User{
			ID: "1",
		}

		subscriptionRepo := mock.NewMockSubscriptionRepositorier(ctrl)
		subscriptionService := NewSubscriptionService(subscriptionRepo)
		userSubscription, err := subscriptionService.CreateUserSubscription(context.WithValue(context.Background(), entity.UserContextKey, &user), dto.CreateUserSubscriptionParam{
			SubscriptionID: "invalid",
		})
		assert.Nil(t, userSubscription)
		assert.NotNil(t, err)
		assert.True(t, errs.IsEqualType(err, errs.ErrUnprocessable))
	})

	t.Run("failed when subscriptio id is free", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		user := entity.User{
			ID: "1",
		}

		subscriptionRepo := mock.NewMockSubscriptionRepositorier(ctrl)
		subscriptionService := NewSubscriptionService(subscriptionRepo)
		userSubscription, err := subscriptionService.CreateUserSubscription(context.WithValue(context.Background(), entity.UserContextKey, &user), dto.CreateUserSubscriptionParam{
			SubscriptionID: "Free",
		})
		assert.Nil(t, userSubscription)
		assert.NotNil(t, err)
		assert.True(t, errs.IsEqualType(err, errs.ErrUnprocessable))
	})

	t.Run("failed to get active user subscription by user id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		user := entity.User{
			ID: "1",
		}

		subscriptionRepo := mock.NewMockSubscriptionRepositorier(ctrl)
		subscriptionRepo.EXPECT().GetActiveUserSubscriptionByUserID(gomock.Any(), user.ID).Return(nil, errs.ErrInternal).Times(1)

		subscriptionService := NewSubscriptionService(subscriptionRepo)
		userSubscription, err := subscriptionService.CreateUserSubscription(context.WithValue(context.Background(), entity.UserContextKey, &user), dto.CreateUserSubscriptionParam{
			UserID:         user.ID,
			SubscriptionID: "Premium",
		})
		assert.Nil(t, userSubscription)
		assert.NotNil(t, err)
		assert.True(t, errs.IsEqualType(err, errs.ErrInternal))
	})

	t.Run("failed when user already have subscription ", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		user := entity.User{
			ID: "1",
		}

		subscriptionRepo := mock.NewMockSubscriptionRepositorier(ctrl)
		subscriptionRepo.EXPECT().GetActiveUserSubscriptionByUserID(gomock.Any(), user.ID).Return(&entity.UserSubscription{}, nil).Times(1)

		subscriptionService := NewSubscriptionService(subscriptionRepo)
		userSubscription, err := subscriptionService.CreateUserSubscription(context.WithValue(context.Background(), entity.UserContextKey, &user), dto.CreateUserSubscriptionParam{
			UserID:         user.ID,
			SubscriptionID: "Premium",
		})
		assert.Nil(t, userSubscription)
		assert.NotNil(t, err)
		assert.True(t, errs.IsEqualType(err, errs.ErrUnprocessable))
	})

	t.Run("failed to get free subscription for the default value", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		user := entity.User{
			ID: "1",
		}

		subscriptionRepo := mock.NewMockSubscriptionRepositorier(ctrl)
		subscriptionRepo.EXPECT().GetActiveUserSubscriptionByUserID(gomock.Any(), user.ID).Return(nil, errs.ErrNotFound).Times(1)
		subscriptionRepo.EXPECT().GetSubscriptionByID(gomock.Any(), "Free").Return(nil, errs.ErrInternal).Times(1)

		subscriptionService := NewSubscriptionService(subscriptionRepo)
		userSubscription, err := subscriptionService.CreateUserSubscription(context.WithValue(context.Background(), entity.UserContextKey, &user), dto.CreateUserSubscriptionParam{
			UserID:         user.ID,
			SubscriptionID: "Premium",
		})
		assert.Nil(t, userSubscription)
		assert.NotNil(t, err)
		assert.True(t, errs.IsEqualType(err, errs.ErrInternal))
	})

	t.Run("failed to create user subscription", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		user := entity.User{
			ID: "1",
		}

		subscription1 := entity.Subscription{ID: "Free"}
		subscription2 := entity.Subscription{ID: "Premium"}

		subscriptionRepo := mock.NewMockSubscriptionRepositorier(ctrl)
		subscriptionRepo.EXPECT().GetActiveUserSubscriptionByUserID(gomock.Any(), user.ID).Return(nil, errs.ErrNotFound).Times(1)
		subscriptionRepo.EXPECT().GetSubscriptionByID(gomock.Any(), "Free").Return(&subscription1, nil).Times(1)
		subscriptionRepo.EXPECT().GetSubscriptionByID(gomock.Any(), "Premium").Return(&subscription2, nil).Times(1)
		subscriptionRepo.EXPECT().CreateUserSubscription(gomock.Any(), gomock.Any()).Return(errs.ErrInternal).Times(1)

		subscriptionService := NewSubscriptionService(subscriptionRepo)
		userSubscription, err := subscriptionService.CreateUserSubscription(context.WithValue(context.Background(), entity.UserContextKey, &user), dto.CreateUserSubscriptionParam{
			UserID:         user.ID,
			SubscriptionID: "Premium",
		})
		assert.Nil(t, userSubscription)
		assert.NotNil(t, err)
		assert.True(t, errs.IsEqualType(err, errs.ErrInternal))
	})

	t.Run("success to create user subscription", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		user := entity.User{
			ID: "1",
		}

		subscription1 := entity.Subscription{ID: "Free"}
		subscription2 := entity.Subscription{ID: "Premium"}

		subscriptionRepo := mock.NewMockSubscriptionRepositorier(ctrl)
		subscriptionRepo.EXPECT().GetActiveUserSubscriptionByUserID(gomock.Any(), user.ID).Return(nil, errs.ErrNotFound).Times(1)
		subscriptionRepo.EXPECT().GetSubscriptionByID(gomock.Any(), "Free").Return(&subscription1, nil).Times(1)
		subscriptionRepo.EXPECT().GetSubscriptionByID(gomock.Any(), "Premium").Return(&subscription2, nil).Times(1)
		subscriptionRepo.EXPECT().CreateUserSubscription(gomock.Any(), gomock.Any()).Return(nil).Times(1)

		subscriptionService := NewSubscriptionService(subscriptionRepo)
		userSubscription, err := subscriptionService.CreateUserSubscription(context.WithValue(context.Background(), entity.UserContextKey, &user), dto.CreateUserSubscriptionParam{
			UserID:         user.ID,
			SubscriptionID: "Premium",
		})
		assert.NotNil(t, userSubscription)
		assert.Nil(t, err)
	})

}
