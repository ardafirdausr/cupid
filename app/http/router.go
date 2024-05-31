package http

import (
	"com.ardafirdausr.cupid/app/http/handler"
	"github.com/labstack/echo/v4"
)

type httpRouter struct {
	userHandler *handler.UserHandler
}

func newRouter(userHandler *handler.UserHandler) *httpRouter {
	return &httpRouter{
		userHandler: userHandler,
	}
}

func (router *httpRouter) setupRouteOnServer(e *echo.Echo) {
	versionGroup := e.Group("/v1")

	// User routes
	userGroup := versionGroup.Group("/user")
	userGroup.PUT("", router.userHandler.Register)
}
