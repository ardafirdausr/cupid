package http

import (
	"com.ardafirdausr.cupid/app/http/handler"
	"github.com/labstack/echo/v4"
)

type httpRouter struct {
	userHandler *handler.UserHandler
	authHandler *handler.AuthHandler
}

func newRouter(
	userHandler *handler.UserHandler,
	authHandler *handler.AuthHandler,
) *httpRouter {
	return &httpRouter{
		userHandler: userHandler,
		authHandler: authHandler,
	}
}

func (router *httpRouter) setupRouteOnServer(e *echo.Echo) {
	versionGroup := e.Group("/v1")

	// User routes
	userGroup := versionGroup.Group("/users")
	userGroup.PUT("/:id", router.userHandler.Update)

	// Auth routes
	authGroup := versionGroup.Group("/auth")
	authGroup.POST("/register", router.authHandler.Register)
	authGroup.POST("/login", router.authHandler.Login)
}
