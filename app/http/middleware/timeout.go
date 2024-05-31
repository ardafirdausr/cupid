package middleware

import (
	"time"

	"com.ardafirdausr.cupid/internal/pkg/logger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func TimeoutMiddleware(timeout time.Duration) echo.MiddlewareFunc {
	return middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout:      timeout,
		ErrorMessage: "Request Timeout",
		OnTimeoutRouteErrorHandler: func(err error, ctx echo.Context) {
			logger.Log.Error().
				Str("method", ctx.Request().Method).
				Str("uri", ctx.Request().RequestURI).
				Msg("REQUEST TIMEOUT")
		},
	})
}
