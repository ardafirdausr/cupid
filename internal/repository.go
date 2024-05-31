package internal

import (
	"context"

	"com.ardafirdausr.cupid/internal/entity"
)

type UserRepositorier interface {
	CreateUser(ctx context.Context, user *entity.User) error
}
