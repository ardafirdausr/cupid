package http

import (
	"com.ardafirdausr.cupid/app/http/handler"
	"com.ardafirdausr.cupid/app/http/middleware"
	"com.ardafirdausr.cupid/internal/entity"
	"github.com/labstack/echo/v4"
)

type httpRouter struct {
	config              entity.CommonConfig
	userHandler         *handler.UserHandler
	authHandler         *handler.AuthHandler
	matchingHandler     *handler.MatchingHandler
	subscriptionHandler *handler.SubscriptionHandler
}

func newRouter(
	config entity.CommonConfig,
	userHandler *handler.UserHandler,
	authHandler *handler.AuthHandler,
	matchingHandler *handler.MatchingHandler,
	subscriptionHandler *handler.SubscriptionHandler,
) *httpRouter {
	return &httpRouter{
		config:              config,
		userHandler:         userHandler,
		authHandler:         authHandler,
		matchingHandler:     matchingHandler,
		subscriptionHandler: subscriptionHandler,
	}
}

func (router *httpRouter) setupRouteOnServer(e *echo.Echo) {
	versionGroup := e.Group("/v1")

	// User routes
	userGroup := versionGroup.Group("/users")
	userGroup.Use(middleware.JWTMiddleware(router.config.JWTSecretKey))
	userGroup.PUT("/:ID", router.userHandler.Update)

	// User subscription routes
	userSubscriptionGroup := userGroup.Group(":userID/subscriptions")
	userSubscriptionGroup.POST("", router.subscriptionHandler.CreateUserSubscription)

	// Auth routes
	authGroup := versionGroup.Group("/auth")
	authGroup.POST("/register", router.authHandler.Register)
	authGroup.POST("/login", router.authHandler.Login)

	// Auth routes
	matchingGroup := versionGroup.Group("/matchings")
	matchingGroup.Use(middleware.JWTMiddleware(router.config.JWTSecretKey))
	matchingGroup.GET("", router.matchingHandler.GetMatchingRecommendations)
	matchingGroup.POST("", router.matchingHandler.CreateMatching)

	// Subscription plan routes
	subscriptionGroup := versionGroup.Group("/subscriptions")
	subscriptionGroup.GET("", router.subscriptionHandler.GetSubscriptions)
}
