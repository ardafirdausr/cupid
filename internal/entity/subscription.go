package entity

import "time"

const (
	SubscriptionFree    SubscriptionID = "Free"
	SubscriptionPremium SubscriptionID = "Premium"
)

type SubscriptionID string

func (id SubscriptionID) String() string {
	return string(id)
}

func (id SubscriptionID) Valid() bool {
	switch id {
	case SubscriptionFree, SubscriptionPremium:
		return true
	}

	return false
}

type SubscriptionFeature struct {
	MaxSwipe int  `json:"max_swipe" bson:"maxSwipe"`
	HasBadge bool `json:"has_badge" bson:"hasBadge"`
}

func (feature *SubscriptionFeature) Merge(key string, reference SubscriptionFeature) bool {
	switch key {
	case "max_swipe":
		feature.MaxSwipe = reference.MaxSwipe
	case "has_badge":
		feature.HasBadge = reference.HasBadge
	}

	return false
}

type Subscription struct {
	ID             SubscriptionID      `json:"id" bson:"_id"`
	Price          float64             `json:"price" bson:"price"`
	DurationInDays int                 `json:"duration_in_days" bson:"durationInDays"`
	Features       SubscriptionFeature `json:"features" bson:"features"`
}

type UserSubscription struct {
	ID                  string              `json:"id" bson:"_id"`
	UserID              string              `json:"user_id" bson:"userID"`
	SubscriptionID      string              `json:"subscription_id" bson:"subscriptionID"`
	SubscriptionFeature SubscriptionFeature `json:"subscription_feature" bson:"subscriptionFeature"`
	PaymentCode         string              `json:"paid_code" bson:"paidCode"`
	PaidAt              time.Time           `json:"paid_at" bson:"paidAt"`
	ExpiredAt           time.Time           `json:"expired_at" bson:"expiredAt"`
}
