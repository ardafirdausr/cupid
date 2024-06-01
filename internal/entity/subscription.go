package entity

import "time"

type SubscriptionID string

const (
	SubscriptionTypeFree    SubscriptionID = "Free"
	SubscriptionTypePremium SubscriptionID = "Premium"
)

type SubscriptionFeature struct {
	MaxSwipe int  `json:"max_swipe" bson:"maxSwipe"`
	HasBadge bool `json:"has_badge" bson:"hasBadge"`
}

type Subscription struct {
	ID             SubscriptionID      `json:"id" bson:"_id"`
	Price          float64             `json:"price" bson:"price"`
	DurationInDays int                 `json:"duration_in_days" bson:"durationInDays"`
	Features       SubscriptionFeature `json:"features" bson:"features"`
}

type UserSubsrciption struct {
	ID                  string    `json:"id" bson:"_id"`
	PaidAt              time.Time `json:"paid_at" bson:"paid_at"`
	SubscriptionID      string    `json:"subscription_id" bson:"subscriptionID"`
	SubscriptionFeature SubscriptionFeature
}
