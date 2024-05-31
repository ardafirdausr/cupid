package handler

import (
	"net/http"

	"com.ardafirdausr.cupid/app/http/handler/response"
	"com.ardafirdausr.cupid/internal"
	"com.ardafirdausr.cupid/internal/dto"
	customJWT "com.ardafirdausr.cupid/internal/pkg/jwt"
	"github.com/labstack/echo/v4"
)

type MatchingHandler struct {
	matchingService internal.MatchingServicer
	jwtHelper       *customJWT.Helper
}

func NewMatchingHandler(matchingService internal.MatchingServicer, jwtHelper *customJWT.Helper) *MatchingHandler {
	return &MatchingHandler{
		matchingService: matchingService,
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

	user, err := customJWT.GetUserFromContext(reqCtx)
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
