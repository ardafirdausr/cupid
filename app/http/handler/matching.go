package handler

import (
	"net/http"

	"com.ardafirdausr.cupid/app/http/handler/response"
	"com.ardafirdausr.cupid/internal"
	"com.ardafirdausr.cupid/internal/dto"
	helper "com.ardafirdausr.cupid/internal/helper"
	"com.ardafirdausr.cupid/internal/pkg/logger"
	"com.ardafirdausr.cupid/internal/pkg/validator"
	"github.com/labstack/echo/v4"
)

type MatchingHandler struct {
	matchingService internal.MatchingServicer
	validator       validator.Validator
	jwtHelper       *helper.ContextInjector
}

func NewMatchingHandler(matchingService internal.MatchingServicer, validator validator.Validator, jwtHelper *helper.ContextInjector) *MatchingHandler {
	return &MatchingHandler{
		matchingService: matchingService,
		validator:       validator,
		jwtHelper:       jwtHelper,
	}
}

func (handler *MatchingHandler) GetMatchingRecommendations(ctx echo.Context) error {
	reqCtx, err := handler.jwtHelper.InjectUserFromJWT(ctx.Request().Context(), ctx.Get("user"))
	if err != nil {
		return err
	}

	var param dto.MatchingRecommendationsFilter
	if err := ctx.Bind(&param); err != nil {
		return err
	}

	user, err := helper.GetUserFromContext(reqCtx)
	if err != nil {
		return err
	}

	param.SetDefault()
	param.SetByUser(user)
	users, err := handler.matchingService.GetMatchingRecommendations(reqCtx, param)
	if err != nil {
		return err
	}

	usersResp := make([]response.UserResponse, 0, len(users))
	for _, user := range users {
		var response response.UserResponse
		response.FromUser(&user)
		usersResp = append(usersResp, response)
	}

	resp := response.BasicResponse{
		Message: "success",
		Data:    usersResp,
	}
	return ctx.JSON(http.StatusOK, resp)

}

func (handler *MatchingHandler) CreateMatching(ctx echo.Context) error {
	reqCtx, err := handler.jwtHelper.InjectUserFromJWT(ctx.Request().Context(), ctx.Get("user"))
	if err != nil {
		return err
	}

	var param dto.CreateMatchingParam
	if err := ctx.Bind(&param); err != nil {
		return err
	}

	if mapErr, err := handler.validator.ValidateStruct(param); err != nil {
		logger.Log.Err(err).Msg("failed to validate request body")
		return ctx.JSON(http.StatusBadRequest, response.BasicErrorResponse{Message: "invalid request body", Errors: mapErr})
	}

	matching, err := handler.matchingService.MatchMaking(reqCtx, param)
	if err != nil {
		return err
	}

	resp := response.BasicResponse{
		Message: "success",
		Data:    matching,
	}
	return ctx.JSON(http.StatusCreated, resp)

}
