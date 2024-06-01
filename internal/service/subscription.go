package service

import (
	"context"
	"time"

	"com.ardafirdausr.cupid/internal"
	"com.ardafirdausr.cupid/internal/dto"
	"com.ardafirdausr.cupid/internal/entity"
	"com.ardafirdausr.cupid/internal/entity/errs"
	"com.ardafirdausr.cupid/internal/helper"
	"github.com/pkg/errors"
)

type SubscriptionServicer struct {
	subscriptionRepo internal.SubscriptionRepositorier
}

func NewSubscriptionService(subscriptionRepo internal.SubscriptionRepositorier) *SubscriptionServicer {
	return &SubscriptionServicer{subscriptionRepo: subscriptionRepo}
}

func (service *SubscriptionServicer) GetAllSubscriptions(ctx context.Context) ([]entity.Subscription, error) {
	subscriptions, err := service.subscriptionRepo.GetAllSubscriptions(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get all subscriptions")
	}

	return subscriptions, nil
}

func (service *SubscriptionServicer) GetActiveUserSubscription(ctx context.Context) (*entity.UserSubscription, error) {
	user, err := helper.GetUserFromContext(ctx)
	if err != nil {
		return nil, errs.NewErrInternal("failed to get user from context")
	}

	// Get existing user subscription
	existingUserSubscription, err := service.subscriptionRepo.GetActiveUserSubscriptionByUserID(ctx, user.ID)
	if err != nil && !errs.IsEqualType(err, errs.ErrNotFound) {
		return nil, errors.Wrap(err, "failed to get active user subscription by user id")
	}

	if existingUserSubscription != nil {
		return existingUserSubscription, nil
	}

	// Get free subscription as default
	freeSubscription, err := service.subscriptionRepo.GetSubscriptionByID(ctx, entity.SubscriptionFree.String())
	if err != nil {
		return nil, errors.Wrap(err, "failed to get subscription by id")
	}

	// Create free subscription for user
	var userSubscription entity.UserSubscription
	userSubscription.UserID = user.ID
	userSubscription.SubscriptionID = freeSubscription.ID.String()
	userSubscription.SubscriptionFeature = freeSubscription.Features

	return &userSubscription, nil
}

func (service *SubscriptionServicer) CreateUserSubscription(ctx context.Context, param dto.CreateUserSubscriptionParam) (*entity.UserSubscription, error) {
	user, err := helper.GetUserFromContext(ctx)
	if err != nil {
		return nil, errs.NewErrInternal("failed to get user from context")
	}

	var userSubscription entity.UserSubscription
	param.ToUserSubscription(&userSubscription)

	if !entity.SubscriptionID(param.SubscriptionID).Valid() {
		return nil, errs.NewErrUnprocessable("invalid subscription id")
	} else if entity.SubscriptionID(param.SubscriptionID) == entity.SubscriptionFree {
		return nil, errs.NewErrUnprocessable("cannot subscribe to free subscription")
	}

	// Get existing user subscription
	existingUserSubscription, err := service.subscriptionRepo.GetActiveUserSubscriptionByUserID(ctx, param.UserID)
	if err != nil && !errs.IsEqualType(err, errs.ErrNotFound) {
		return nil, errors.Wrap(err, "failed to get active user subscription by user id")
	} else if existingUserSubscription != nil {
		return nil, errs.NewErrUnprocessable("user already has active subscription")
	}

	// Get free subscription as default
	freeSubscription, err := service.subscriptionRepo.GetSubscriptionByID(ctx, entity.SubscriptionFree.String())
	if err != nil {
		return nil, errors.Wrap(err, "failed to get subscription by id")
	}

	// Get detail of subscription
	newSubscription, err := service.subscriptionRepo.GetSubscriptionByID(ctx, param.SubscriptionID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get subscription by id")
	}

	now := time.Now()
	userSubscription.UserID = user.ID
	userSubscription.ExpiredAt = now.AddDate(0, 0, newSubscription.DurationInDays)
	userSubscription.SubscriptionFeature = freeSubscription.Features
	for _, feature := range param.SubscriptionFeature {
		userSubscription.SubscriptionFeature.Merge(feature, newSubscription.Features)
	}

	if err := service.subscriptionRepo.CreateUserSubscription(ctx, &userSubscription); err != nil {
		return nil, errors.Wrap(err, "failed to create user subscription")
	}

	return &userSubscription, nil
}
