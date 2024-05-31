package handler

import (
	"net/http"

	"com.ardafirdausr.cupid/internal"
	"com.ardafirdausr.cupid/internal/dto"
	"com.ardafirdausr.cupid/internal/pkg/helper"
	"github.com/labstack/echo/v4"
)

type MatchingHandler struct {
	matchingService internal.MatchingServicer
	injector        *helper.Injector
}

func NewMatchingHandler(matchingService internal.MatchingServicer, injector *helper.Injector) *MatchingHandler {
	return &MatchingHandler{
		matchingService: matchingService,
		injector:        injector,
	}
}

func (handler *MatchingHandler) GetMatchingRecommendations(ctx echo.Context) error {
	reqCtx, err := handler.injector.InjectUserFromJwt(ctx.Request().Context())
	if err != nil {
		return err
	}

	var param dto.MatchingRecommendationsFilter
	data, err := handler.matchingService.GetMatchingRecommendations(reqCtx, param)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, data)

}
