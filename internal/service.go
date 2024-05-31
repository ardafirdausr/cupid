//go:generate mockgen -source service.go -package mock -destination ./mock/service.go
package internal

import (
	"context"

	"com.ardafirdausr.cupid/internal/dto"
	"com.ardafirdausr.cupid/internal/entity"
)

type UserServicer interface {
	UpdateUser(ctx context.Context, param dto.UpdateUserParam) (*entity.User, error)
}

type AuthServicer interface {
	Register(ctx context.Context, param dto.RegisterUserParam) (*entity.User, string, error)
	Login(ctx context.Context, param dto.LoginrUserParam) (*entity.User, string, error)
}
