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

type UserHandler struct {
	userService internal.UserServicer
	validator   validator.Validator
}

func NewUserHandler(userService internal.UserServicer, validator validator.Validator) *UserHandler {
	return &UserHandler{
		userService: userService,
		validator:   validator,
	}
}

func (handler *UserHandler) Update(ctx echo.Context) error {
	var param dto.UpdateUserParam
	if err := ctx.Bind(&param); err != nil {
		logger.Log.Err(err).Msg("failed to bind request body")
		return errs.NewErrInvalidData("invalid request body")
	}

	param.ID = ctx.Param("id")
	if mapErr, err := handler.validator.ValidateStruct(param); err != nil {
		logger.Log.Err(err).Msg("failed to validate request body")
		return ctx.JSON(http.StatusBadRequest, response.BasicErrorResponse{Message: "invalid request body", Errors: mapErr})
	}

	user, err := handler.userService.UpdateUser(ctx.Request().Context(), param)
	if err != nil {
		return err
	}

	var userResp response.UserResponse
	userResp.FromUser(user)
	resp := response.BasicResponse{Message: "User registered", Data: userResp}
	return ctx.JSON(http.StatusCreated, resp)
}
