package helper

import (
	"context"
	"errors"

	"com.ardafirdausr.cupid/internal"
	"com.ardafirdausr.cupid/internal/entity"
	"github.com/golang-jwt/jwt/v5"
)

type Injector struct {
	userService internal.UserServicer
}

func Newinjector(userService internal.UserServicer) *Injector {
	return &Injector{userService: userService}
}

func (inj *Injector) InjectUserFromJwt(ctx context.Context) (context.Context, error) {
	userJWT := ctx.Value("user")
	if userJWT == nil {
		return ctx, errors.New("user not found")
	}

	if userJWT, ok := userJWT.(*jwt.Token); ok {
		claims := userJWT.Claims.(jwt.MapClaims)
		userID := claims["id"].(string)

		user, err := inj.userService.GetUserByID(ctx, userID)
		if err != nil {
			return ctx, err
		}

		return context.WithValue(ctx, entity.UserContextKey, user), nil
	}

	return ctx, nil
}
