package handler

import (
	"net/http"

	"com.ardafirdausr.cupid/app/http/handler/response"
	"com.ardafirdausr.cupid/internal"
	"com.ardafirdausr.cupid/internal/dto"
	"com.ardafirdausr.cupid/internal/entity/errs"
	"com.ardafirdausr.cupid/internal/pkg/logger"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService internal.UserServicer
}

func NewUserHandler(userService internal.UserServicer) *UserHandler {
	return &UserHandler{userService: userService}
}

func (handler *UserHandler) Register(ctx echo.Context) error {
	var param dto.RegisterUserParam
	if err := ctx.Bind(&param); err != nil {
		logger.Log.Err(err).Msg("failed to bind request body")
		return errs.NewErrInvalidData("invalid request body")
	}

	user, err := handler.userService.RegisterUser(ctx.Request().Context(), param)
	if err != nil {
		return err
	}

	res := response.BasicResponse{Message: "User registered", Data: user}
	return ctx.JSON(http.StatusCreated, res)
}
