package helper

import (
	"context"
	"errors"

	"com.ardafirdausr.cupid/internal"
	"com.ardafirdausr.cupid/internal/entity"
	"com.ardafirdausr.cupid/internal/entity/errs"
	"com.ardafirdausr.cupid/internal/pkg/logger"
	"github.com/golang-jwt/jwt/v5"
)

type ContextInjector struct {
	userService         internal.UserServicer
	subscriptionService internal.SubscriptionServicer
}

func NewContextInjector(
	userService internal.UserServicer,
	subscriptionService internal.SubscriptionServicer,
) *ContextInjector {
	return &ContextInjector{
		userService:         userService,
		subscriptionService: subscriptionService,
	}
}

func (inj *ContextInjector) InjectUserFromJWT(ctx context.Context, jwtToken interface{}) (context.Context, error) {
	userJWT, ok := jwtToken.(*jwt.Token)
	if !ok {
		logger.Log.Err(errors.New("invalid jwt token")).Msg("failed to get user from jwt token")
		return ctx, errs.NewErrUnauthorized("invalid jwt token")
	}

	userID, err := userJWT.Claims.GetSubject()
	if err != nil {
		logger.Log.Err(err).Msg("failed to get user id from jwt token")
		return ctx, errs.NewErrUnauthorized("invalid jwt token")
	}

	user, err := inj.userService.GetUserByID(ctx, userID)
	if err != nil {
		return ctx, err
	}

	ctx = context.WithValue(ctx, entity.UserContextKey, user)

	userSubscription, err := inj.subscriptionService.GetActiveUserSubscription(ctx)
	if err != nil {
		return ctx, err
	}

	ctx = context.WithValue(ctx, entity.UserSubscriptionContextKey, userSubscription)
	return ctx, nil
}

func GetUserFromContext(ctx context.Context) (*entity.User, error) {
	user, ok := ctx.Value(entity.UserContextKey).(*entity.User)
	if !ok {
		return nil, errs.NewErrUnauthorized("user not found in context")
	}

	return user, nil
}

func GetSubscriptionFromContext(ctx context.Context) (*entity.UserSubscription, error) {
	user, ok := ctx.Value(entity.UserSubscriptionContextKey).(*entity.UserSubscription)
	if !ok {
		return nil, errs.NewErrUnauthorized("user not found in context")
	}

	return user, nil
}
