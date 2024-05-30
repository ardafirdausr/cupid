package middleware

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
)

func TimeoutMiddleware(timeoutInMilis int, logger *zerolog.Logger) echo.MiddlewareFunc {
	if timeoutInMilis <= 1000 {
		timeoutInMilis = 10000
	}

	return middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout:      time.Duration(timeoutInMilis) * time.Millisecond,
		ErrorMessage: "Request Timeout",
		OnTimeoutRouteErrorHandler: func(err error, ctx echo.Context) {
			logger.Error().
				Str("method", ctx.Request().Method).
				Str("uri", ctx.Request().RequestURI).
				Msg("REQUEST TIMEOUT")
		},
	})
}
