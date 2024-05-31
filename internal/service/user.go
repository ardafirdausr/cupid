package service

import (
	"context"

	"com.ardafirdausr.cupid/internal"
	"com.ardafirdausr.cupid/internal/dto"
	"com.ardafirdausr.cupid/internal/entity"
	"github.com/pkg/errors"
)

type UserService struct {
	userRepo internal.UserRepositorier
}

func NewUserService(userRepo internal.UserRepositorier) *UserService {
	return &UserService{userRepo: userRepo}
}

func (svc *UserService) RegisterUser(ctx context.Context, param dto.RegisterUserParam) (*entity.User, error) {
	var user entity.User
	param.ToUser(&user)

	if err := svc.userRepo.CreateUser(ctx, &user); err != nil {
		return nil, errors.Wrap(err, "failed to create user")
	}

	return &user, nil
}
