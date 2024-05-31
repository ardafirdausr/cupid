package http

import (
	"com.ardafirdausr.cupid/app/http/handler"
	"com.ardafirdausr.cupid/app/http/middleware"
	"com.ardafirdausr.cupid/internal/entity"
	"github.com/labstack/echo/v4"
)

type httpRouter struct {
	config          entity.CommonConfig
	userHandler     *handler.UserHandler
	authHandler     *handler.AuthHandler
	matchingHandler *handler.MatchingHandler
}

func newRouter(
	config entity.CommonConfig,
	userHandler *handler.UserHandler,
	authHandler *handler.AuthHandler,
	matchingHandler *handler.MatchingHandler,
) *httpRouter {
	return &httpRouter{
		config:          config,
		userHandler:     userHandler,
		authHandler:     authHandler,
		matchingHandler: matchingHandler,
	}
}

func (router *httpRouter) setupRouteOnServer(e *echo.Echo) {
	versionGroup := e.Group("/v1")

	// User routes
	userGroup := versionGroup.Group("/users")
	userGroup.Use(middleware.JWTMiddleware(router.config.JWTSecretKey))
	userGroup.PUT("/:id", router.userHandler.Update)

	// Auth routes
	authGroup := versionGroup.Group("/auth")
	authGroup.POST("/register", router.authHandler.Register)
	authGroup.POST("/login", router.authHandler.Login)

	// Auth routes
	matchingGroup := versionGroup.Group("/matchings")
	matchingGroup.Use(middleware.JWTMiddleware(router.config.JWTSecretKey))
	matchingGroup.GET("", router.matchingHandler.GetMatchingRecommendations)

}
