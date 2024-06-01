package service

import (
	"context"

	"com.ardafirdausr.cupid/internal"
	"com.ardafirdausr.cupid/internal/dto"
	"com.ardafirdausr.cupid/internal/entity"
	"com.ardafirdausr.cupid/internal/entity/errs"
	"github.com/pkg/errors"
)

type UserService struct {
	userRepo internal.UserRepositorier
}

func NewUserService(userRepo internal.UserRepositorier) *UserService {
	return &UserService{userRepo: userRepo}
}

func (svc *UserService) GetUserByID(ctx context.Context, id string) (*entity.User, error) {
	user, err := svc.userRepo.GetUserByID(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user by ID")
	}

	return user, nil
}

func (svc *UserService) UpdateUser(ctx context.Context, param dto.UpdateUserParam) (*entity.User, error) {
	user, err := svc.userRepo.GetUserByID(ctx, param.ID)
	if err != nil && !errs.IsEqualType(err, errs.ErrNotFound) {
		return nil, errors.Wrap(err, "failed to get user by ID")
	}

	if user == nil {
		return nil, errs.NewErrNotFound("user not found")
	}

	param.ToUser(user)
	if err := svc.userRepo.UpdateUserByID(ctx, param.ID, user); err != nil {
		return nil, errors.Wrap(err, "failed to update user")
	}

	return user, nil
}
