package handler

import (
	"net/http"

	"com.ardafirdausr.cupid/app/http/handler/response"
	"com.ardafirdausr.cupid/internal"
	"com.ardafirdausr.cupid/internal/dto"
	"com.ardafirdausr.cupid/internal/entity/errs"
	helper "com.ardafirdausr.cupid/internal/helper"
	"com.ardafirdausr.cupid/internal/pkg/logger"
	"com.ardafirdausr.cupid/internal/pkg/validator"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService internal.UserServicer
	jwtHelper   *helper.ContextInjector
	validator   validator.Validator
}

func NewUserHandler(
	userService internal.UserServicer,
	jwtHelper *helper.ContextInjector,
	validator validator.Validator,
) *UserHandler {
	return &UserHandler{
		userService: userService,
		validator:   validator,
	}
}

func (handler *UserHandler) Update(ctx echo.Context) error {
	reqCtx, err := handler.jwtHelper.InjectUserFromJWT(ctx.Request().Context(), ctx.Get("user"))
	if err != nil {
		return err
	}

	reqUser, err := helper.GetUserFromContext(reqCtx)
	if err != nil {
		return err
	}

	var param dto.UpdateUserParam
	if err := ctx.Bind(&param); err != nil {
		logger.Log.Err(err).Msg("failed to bind request body")
		return errs.NewErrInvalidData("invalid request body")
	}

	param.ID = ctx.Param("ID")
	if param.ID != reqUser.ID {
		return errs.NewErrForbidden("forbidden to update other user")
	}

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
