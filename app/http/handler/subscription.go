package handler

import (
	"com.ardafirdausr.cupid/app/http/handler/response"
	"com.ardafirdausr.cupid/internal"
	"github.com/labstack/echo/v4"
)

type SubscriptionHandler struct {
	subscriptionService internal.SubscriptionServicer
}

func NewSubscriptionHandler(subscriptionService internal.SubscriptionServicer) *SubscriptionHandler {
	return &SubscriptionHandler{subscriptionService: subscriptionService}
}

func (handler *SubscriptionHandler) GetSubscriptions(ctx echo.Context) error {
	subscriptions, err := handler.subscriptionService.GetAllSubscriptions(ctx.Request().Context())
	if err != nil {
		return err
	}

	return ctx.JSON(200, response.BasicResponse{
		Message: "success",
		Data:    subscriptions,
	})
}
