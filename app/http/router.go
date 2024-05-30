package http

import (
	"com.ardafirdausr.cupid/app/http/handler"
	"github.com/labstack/echo/v4"
)

type httpRouter struct{}

func newRouter() *httpRouter {
	return &httpRouter{}
}

func (router *httpRouter) setupRouteOnServer(e *echo.Echo) {

	// hello
	helloHandler := handler.NewHelloHandler()
	e.GET("/hello", helloHandler.Hello)
}
