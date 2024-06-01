package dto

import (
	"time"

	"com.ardafirdausr.cupid/internal/entity"
)

type CreateUserSubscriptionParam struct {
	UserID              string   `path:"userID" validate:"required"`
	PaymentCode         string   `json:"payment_code" validate:"required"`
	SubscriptionID      string   `json:"subscription_id" validate:"required"`
	SubscriptionFeature []string `json:"subscription_feature" validate:"required,max=1"`
}

func (param *CreateUserSubscriptionParam) ToUserSubscription(userSubscription *entity.UserSubscription) {
	userSubscription.PaymentCode = param.PaymentCode
	userSubscription.PaidAt = time.Now()
	userSubscription.SubscriptionID = param.SubscriptionID
}
