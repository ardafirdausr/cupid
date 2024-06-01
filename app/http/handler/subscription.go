package handler

import (
	"net/http"

	"com.ardafirdausr.cupid/app/http/handler/response"
	"com.ardafirdausr.cupid/internal"
	"com.ardafirdausr.cupid/internal/dto"
	"com.ardafirdausr.cupid/internal/entity/errs"
	helper "com.ardafirdausr.cupid/internal/helper"
	"com.ardafirdausr.cupid/internal/pkg/validator"
	"github.com/labstack/echo/v4"
)

type SubscriptionHandler struct {
	subscriptionService internal.SubscriptionServicer
	validator           validator.Validator
	jwtHelper           *helper.ContextInjector
}

func NewSubscriptionHandler(
	subscriptionService internal.SubscriptionServicer,
	validator validator.Validator,
	jwtHelper *helper.ContextInjector,
) *SubscriptionHandler {
	return &SubscriptionHandler{
		subscriptionService: subscriptionService,
		validator:           validator,
		jwtHelper:           jwtHelper,
	}
}

func (handler *SubscriptionHandler) GetSubscriptions(ctx echo.Context) error {
	subscriptions, err := handler.subscriptionService.GetAllSubscriptions(ctx.Request().Context())
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, response.BasicResponse{
		Message: "success",
		Data:    subscriptions,
	})
}

func (handler *SubscriptionHandler) CreateUserSubscription(ctx echo.Context) error {
	reqCtx, err := handler.jwtHelper.InjectUserFromJWT(ctx.Request().Context(), ctx.Get("user"))
	if err != nil {
		return err
	}

	reqUser, err := helper.GetUserFromContext(reqCtx)
	if err != nil {
		return err
	}

	var param dto.CreateUserSubscriptionParam
	if err := ctx.Bind(&param); err != nil {
		return err
	}

	param.UserID = ctx.Param("userID")
	if mapErr, err := handler.validator.ValidateStruct(param); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.BasicErrorResponse{Message: "invalid request body", Errors: mapErr})
	}

	if param.UserID != reqUser.ID {
		return errs.NewErrForbidden("forbidden to update other user")
	}

	userSubscription, err := handler.subscriptionService.CreateUserSubscription(reqCtx, param)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, response.BasicResponse{
		Message: "success",
		Data:    userSubscription,
	})
}

func (handler *SubscriptionHandler) GetUserSubscription(ctx echo.Context) error {
	reqCtx, err := handler.jwtHelper.InjectUserFromJWT(ctx.Request().Context(), ctx.Get("user"))
	if err != nil {
		return err
	}

	userSubscription, err := handler.subscriptionService.GetActiveUserSubscription(reqCtx)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, response.BasicResponse{
		Message: "success",
		Data:    userSubscription,
	})
}
