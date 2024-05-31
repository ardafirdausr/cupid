package middleware

import (
	echoJWT "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware(secret string) echo.MiddlewareFunc {
	return echoJWT.WithConfig(echoJWT.Config{
		SigningKey: []byte(secret),
	})
}
