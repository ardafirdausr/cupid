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

func TestRegister(t *testing.T) {
	t.Run("failed when getting user by email", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		config := entity.CommonConfig{}
		userRepo := mock.NewMockUserRepositorier(ctrl)
		userRepo.EXPECT().GetUserByEmail(gomock.Any(), gomock.Any()).Return(nil, errs.ErrInternal).Times(1)

		param := dto.RegisterUserParam{}
		userService := NewAuthService(config, userRepo)
		user, token, err := userService.Register(context.Background(), param)
		assert.Nil(t, user)
		assert.Equal(t, token, "")
		assert.NotNil(t, err)
		assert.True(t, errs.IsEqualType(err, errs.ErrInternal))
	})

	t.Run("failed when email already registered", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		config := entity.CommonConfig{}
		userRepo := mock.NewMockUserRepositorier(ctrl)
		userRepo.EXPECT().GetUserByEmail(gomock.Any(), gomock.Any()).Return(&entity.User{}, nil).Times(1)

		param := dto.RegisterUserParam{}
		userService := NewAuthService(config, userRepo)
		user, token, err := userService.Register(context.Background(), param)
		assert.Nil(t, user)
		assert.Equal(t, token, "")
		assert.NotNil(t, err)
		assert.True(t, errs.IsEqualType(err, errs.ErrUnprocessable))
	})

	t.Run("failed when creating user", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		config := entity.CommonConfig{}
		userRepo := mock.NewMockUserRepositorier(ctrl)
		userRepo.EXPECT().GetUserByEmail(gomock.Any(), gomock.Any()).Return(nil, errs.ErrNotFound).Times(1)
		userRepo.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(errs.ErrInternal).Times(1)

		param := dto.RegisterUserParam{}
		userService := NewAuthService(config, userRepo)
		user, token, err := userService.Register(context.Background(), param)
		assert.Nil(t, user)
		assert.Equal(t, token, "")
		assert.NotNil(t, err)
		assert.True(t, errs.IsEqualType(err, errs.ErrInternal))
	})

	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		config := entity.CommonConfig{}
		userRepo := mock.NewMockUserRepositorier(ctrl)
		userRepo.EXPECT().GetUserByEmail(gomock.Any(), gomock.Any()).Return(nil, errs.ErrNotFound).Times(1)
		userRepo.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(nil).Times(1)

		param := dto.RegisterUserParam{}
		userService := NewAuthService(config, userRepo)
		user, token, err := userService.Register(context.Background(), param)
		assert.NotNil(t, user)
		assert.NotEmpty(t, token)
		assert.Nil(t, err)
	})
}

func TestLogin(t *testing.T) {
	t.Run("failed when getting user by email", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		config := entity.CommonConfig{}
		userRepo := mock.NewMockUserRepositorier(ctrl)
		userRepo.EXPECT().GetUserByEmail(gomock.Any(), gomock.Any()).Return(nil, errs.ErrInternal).Times(1)

		param := dto.LoginUserParam{}
		userService := NewAuthService(config, userRepo)
		user, token, err := userService.Login(context.Background(), param)
		assert.Nil(t, user)
		assert.Equal(t, token, "")
		assert.NotNil(t, err)
		assert.True(t, errs.IsEqualType(err, errs.ErrInternal))
	})

	t.Run("failed when user not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		config := entity.CommonConfig{}
		userRepo := mock.NewMockUserRepositorier(ctrl)
		userRepo.EXPECT().GetUserByEmail(gomock.Any(), gomock.Any()).Return(nil, nil).Times(1)

		param := dto.LoginUserParam{}
		userService := NewAuthService(config, userRepo)
		user, token, err := userService.Login(context.Background(), param)
		assert.Nil(t, user)
		assert.Equal(t, token, "")
		assert.NotNil(t, err)
		assert.True(t, errs.IsEqualType(err, errs.ErrUnprocessable))
	})

	t.Run("failed when password incorrect", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		config := entity.CommonConfig{}
		userRepo := mock.NewMockUserRepositorier(ctrl)
		userRepo.EXPECT().GetUserByEmail(gomock.Any(), gomock.Any()).Return(&entity.User{}, nil).Times(1)

		param := dto.LoginUserParam{}
		userService := NewAuthService(config, userRepo)
		user, token, err := userService.Login(context.Background(), param)
		assert.Nil(t, user)
		assert.Equal(t, token, "")
		assert.NotNil(t, err)
		assert.True(t, errs.IsEqualType(err, errs.ErrUnprocessable))
	})

	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		config := entity.CommonConfig{}
		userRepo := mock.NewMockUserRepositorier(ctrl)
		userRepo.EXPECT().GetUserByEmail(gomock.Any(), gomock.Any()).Return(&entity.User{Password: "D5cSyXrs1iu74zzASs2GtjMclUmZR4cbFLFXJTcE+ow="}, nil).Times(1)

		param := dto.LoginUserParam{Password: "Rahasia123"}
		userService := NewAuthService(config, userRepo)
		user, token, err := userService.Login(context.Background(), param)
		assert.NotNil(t, user)
		assert.NotEmpty(t, token)
		assert.Nil(t, err)
	})
}

func TestGenerateAuthToken(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		config := entity.CommonConfig{}
		userRepo := mock.NewMockUserRepositorier(ctrl)

		userService := NewAuthService(config, userRepo)
		token, err := userService.generateAuthToken(&entity.User{})
		assert.Nil(t, err)
		assert.NotEmpty(t, token)
	})
}
