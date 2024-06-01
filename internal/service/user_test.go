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

func TestGetUserByID(t *testing.T) {
	t.Run("failed when getting user data", func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()

		userID := "123"

		userRepo := mock.NewMockUserRepositorier(ctl)
		userRepo.EXPECT().GetUserByID(gomock.Any(), userID).Return(nil, errs.ErrNotFound).Times(1)

		userService := NewUserService(userRepo)
		res, err := userService.GetUserByID(context.Background(), userID)
		assert.Nil(t, res)
		assert.NotNil(t, err)
		assert.True(t, errs.IsEqualType(err, errs.ErrNotFound))
	})

	t.Run("success when getting user data", func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()

		userID := "123"
		userData := &entity.User{ID: userID}

		userRepo := mock.NewMockUserRepositorier(ctl)
		userRepo.EXPECT().GetUserByID(gomock.Any(), userID).Return(userData, nil).Times(1)

		userService := NewUserService(userRepo)
		res, err := userService.GetUserByID(context.Background(), userID)
		assert.Nil(t, err)
		assert.Equal(t, userData, res)
	})

}

func TestUpdateUser(t *testing.T) {
	t.Run("failed when getting user data", func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()

		userID := "123"
		param := dto.UpdateUserParam{ID: userID}

		userRepo := mock.NewMockUserRepositorier(ctl)
		userRepo.EXPECT().GetUserByID(gomock.Any(), userID).Return(nil, errs.ErrNotFound).Times(1)

		userService := NewUserService(userRepo)
		res, err := userService.UpdateUser(context.Background(), param)
		assert.Nil(t, res)
		assert.NotNil(t, err)
		assert.True(t, errs.IsEqualType(err, errs.ErrNotFound))
	})

	t.Run("failed when updating user data", func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()

		userID := "123"
		param := dto.UpdateUserParam{ID: userID}
		userData := &entity.User{ID: userID}

		userRepo := mock.NewMockUserRepositorier(ctl)
		userRepo.EXPECT().GetUserByID(gomock.Any(), userID).Return(userData, nil).Times(1)
		userRepo.EXPECT().UpdateUserByID(gomock.Any(), userID, userData).Return(errs.ErrInternal).Times(1)

		userService := NewUserService(userRepo)
		res, err := userService.UpdateUser(context.Background(), param)
		assert.Nil(t, res)
		assert.NotNil(t, err)
		assert.True(t, errs.IsEqualType(err, errs.ErrInternal))
	})

	t.Run("success when updating user data", func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()

		userID := "123"
		param := dto.UpdateUserParam{ID: userID}
		userData := &entity.User{ID: userID}

		userRepo := mock.NewMockUserRepositorier(ctl)
		userRepo.EXPECT().GetUserByID(gomock.Any(), userID).Return(userData, nil).Times(1)
		userRepo.EXPECT().UpdateUserByID(gomock.Any(), userID, userData).Return(nil).Times(1)

		userService := NewUserService(userRepo)
		res, err := userService.UpdateUser(context.Background(), param)
		assert.Nil(t, err)
		assert.Equal(t, userData, res)
	})
}
