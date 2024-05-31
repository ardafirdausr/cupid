package handler

import (
	"net/http"

	"com.ardafirdausr.cupid/app/http/handler/response"
	"com.ardafirdausr.cupid/internal"
	"com.ardafirdausr.cupid/internal/dto"
	"com.ardafirdausr.cupid/internal/entity/errs"
	"com.ardafirdausr.cupid/internal/pkg/logger"
	"com.ardafirdausr.cupid/internal/pkg/validator"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	userService internal.AuthServicer
	validator   validator.Validator
}

func NewAuthHandler(userService internal.AuthServicer, validator validator.Validator) *AuthHandler {
	return &AuthHandler{
		userService: userService,
		validator:   validator,
	}
}

func (handler *AuthHandler) Register(ctx echo.Context) error {
	var param dto.RegisterUserParam
	if err := ctx.Bind(&param); err != nil {
		logger.Log.Err(err).Msg("failed to bind request body")
		return errs.NewErrInvalidData("invalid request body")
	}

	if mapErr, err := handler.validator.ValidateStruct(param); err != nil {
		logger.Log.Err(err).Msg("failed to validate request body")
		return ctx.JSON(http.StatusBadRequest, response.BasicErrorResponse{Message: "invalid request body", Errors: mapErr})
	}

	user, token, err := handler.userService.Register(ctx.Request().Context(), param)
	if err != nil {
		return err
	}

	var userResp response.UserResponse
	userResp.FromUser(user)
	resp := response.BasicResponse{Message: "Register success", Data: userResp, Meta: map[string]interface{}{"token": token}}
	return ctx.JSON(http.StatusCreated, resp)
}

func (handler *AuthHandler) Login(ctx echo.Context) error {
	var param dto.LoginrUserParam
	if err := ctx.Bind(&param); err != nil {
		logger.Log.Err(err).Msg("failed to bind request body")
		return errs.NewErrInvalidData("invalid request body")
	}

	if mapErr, err := handler.validator.ValidateStruct(param); err != nil {
		logger.Log.Err(err).Msg("failed to validate request body")
		return ctx.JSON(http.StatusBadRequest, response.BasicErrorResponse{Message: "invalid request body", Errors: mapErr})
	}

	user, token, err := handler.userService.Login(ctx.Request().Context(), param)
	if err != nil {
		return err
	}

	var userResp response.UserResponse
	userResp.FromUser(user)
	resp := response.BasicResponse{Message: "Login success", Data: userResp, Meta: map[string]interface{}{"token": token}}
	return ctx.JSON(http.StatusCreated, resp)
}
