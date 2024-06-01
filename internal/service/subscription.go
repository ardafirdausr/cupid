package service

import (
	"context"

	"com.ardafirdausr.cupid/internal"
	"com.ardafirdausr.cupid/internal/entity"
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
