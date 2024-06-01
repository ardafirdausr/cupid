package entity

import "time"

type UserSubsrciption struct {
	ID                 string    `json:"id" bson:"_id"`
	SubscriptionPlanID string    `json:"subscription_plan_id" bson:"subscription_plan_id"`
	PaidAt             time.Time `json:"paid_at" bson:"paid_at"`
	ChoosenFeatures    SubscriptionFeature
}

type SubscriptionPlan struct {
	ID             SubscriptionID      `json:"id" bson:"_id"`
	Price          float64             `json:"price" bson:"price"`
	Features       SubscriptionFeature `json:"features" bson:"features"`
	DurationInDays int                 `json:"duration_in_days" bson:"duration_in_days"`
}

type SubscriptionID string

const (
	SubscriptionTypeFree    SubscriptionID = "Free"
	SubscriptionTypePremium SubscriptionID = "Premium"
)

type SubscriptionFeature map[SubscriptionFeatureKey]interface{}

type SubscriptionFeatureKey string

const (
	SubscriptionFeatureMaxSwipe SubscriptionFeatureKey = "max_swipe"
	SubscriptionFeatureHasBadge SubscriptionFeatureKey = "had_badge"
)
