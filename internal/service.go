package internal

import (
	"context"

	"com.ardafirdausr.cupid/internal/dto"
	"com.ardafirdausr.cupid/internal/entity"
)

type UserServicer interface {
	RegisterUser(ctx context.Context, param dto.RegisterUserParam) (*entity.User, error)
}
