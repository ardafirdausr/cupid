package service

import (
	"context"
	"time"

	"com.ardafirdausr.cupid/internal"
	"com.ardafirdausr.cupid/internal/dto"
	"com.ardafirdausr.cupid/internal/entity"
	"com.ardafirdausr.cupid/internal/entity/errs"
	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
)

type AuthService struct {
	config   entity.CommonConfig
	userRepo internal.UserRepositorier
}

func NewAuthService(config entity.CommonConfig, userRepo internal.UserRepositorier) *AuthService {
	return &AuthService{config: config, userRepo: userRepo}
}

func (svc AuthService) Register(ctx context.Context, param dto.RegisterUserParam) (*entity.User, string, error) {
	user := &entity.User{}
	param.ToUser(user)

	if existingUser, err := svc.userRepo.GetUserByEmail(ctx, user.Email); err != nil && !errs.IsEqualType(err, errs.ErrNotFound) {
		return nil, "", errors.Wrap(err, "failed to get user by email")
	} else if existingUser != nil {
		return nil, "", errs.NewErrUnprocessable("email already registered")
	}

	if err := svc.userRepo.CreateUser(ctx, user); err != nil {
		return nil, "", errors.Wrap(err, "failed to create user")
	}

	token, err := svc.generateAuthToken(user)
	if err != nil {
		return nil, "", errors.Wrap(err, "failed to generate token")
	}

	return user, token, nil
}

func (svc AuthService) Login(ctx context.Context, param dto.LoginrUserParam) (*entity.User, string, error) {
	user, err := svc.userRepo.GetUserByEmail(ctx, param.Email)
	if err != nil && !errs.IsEqualType(err, errs.ErrNotFound) {
		return nil, "", errors.Wrap(err, "failed to get user by email")
	}

	if user == nil {
		return nil, "", errs.NewErrUnprocessable("incorrect email or password")
	}

	if !user.ComparePassword(param.Password) {
		return nil, "", errs.NewErrUnprocessable("incorrect email or password")
	}

	token, err := svc.generateAuthToken(user)
	if err != nil {
		return nil, "", errors.Wrap(err, "failed to generate token")
	}

	return user, token, nil
}

func (svc AuthService) generateAuthToken(user *entity.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   user.ID,
		"iat":   jwt.TimeFunc().Unix(),
		"exp":   jwt.TimeFunc().Add(2 * time.Hour).Unix(),
		"email": user.Email,
	})

	tokenString, err := token.SignedString([]byte(svc.config.JWTSecretKey))
	if err != nil {
		return "", errors.Wrap(err, "failed to sign token")
	}

	return tokenString, nil
}
