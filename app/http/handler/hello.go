package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HelloHandler struct{}

func NewHelloHandler() *HelloHandler {
	return &HelloHandler{}
}

func (handler *HelloHandler) Hello(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello, World!")
}
