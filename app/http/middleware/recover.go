package middleware

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
)

func RecoverMiddleware(loggy *zerolog.Logger) echo.MiddlewareFunc {
	return middleware.RecoverWithConfig(middleware.RecoverConfig{
		Skipper: middleware.DefaultSkipper,
		LogErrorFunc: func(c echo.Context, err error, stack []byte) error {
			loggy.Panic().
				Str("method", c.Request().Method).
				Str("uri", c.Request().RequestURI).
				Msg(fmt.Sprintf("Recovery : %s", err.Error()))
			return err
		},
		StackSize:         4 << 10,
		DisableStackAll:   true,
		DisablePrintStack: false,
	})
}
