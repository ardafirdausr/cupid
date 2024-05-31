package middleware

import (
	"fmt"
	"regexp"
	"strings"

	"com.ardafirdausr.cupid/internal/pkg/logger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func DumpLogMiddleware() echo.MiddlewareFunc {
	return middleware.BodyDumpWithConfig(middleware.BodyDumpConfig{
		Skipper: func(c echo.Context) bool {
			skipResourceAccess := []string{".json", ".yaml", ".html", "document", "pdf", ".js", ".css"}
			for _, resource := range skipResourceAccess {
				if strings.Contains(c.Request().URL.Path, resource) {
					return true
				}
			}

			return false
		},
		Handler: (func(c echo.Context, reqBody, resBody []byte) {
			reg, _ := regexp.Compile(`\s\s+`)

			var logRequest interface{}
			reqBody = reg.ReplaceAll([]byte(reqBody), []byte(" "))
			logRequest = strings.TrimSpace(string(reqBody))

			// IF request contain form data, logRequest set nil, to anticipate print file on log
			if strings.Contains(string(reqBody), "form-data") {
				logRequest = nil
			}

			var logResponse interface{}
			resBody = reg.ReplaceAll([]byte(resBody), []byte(" "))
			logResponse = strings.TrimSpace(string(resBody))

			logMessage := fmt.Sprintf("[HTTP REQUEST] %s %s", c.Request().Method, c.Request().URL)
			errorCode := c.Response().Status
			switch {
			case errorCode >= 200 && errorCode < 300:
				logger.Log.Info().
					Interface("request", logRequest).
					Interface("response", logResponse).
					Msg(logMessage)
			case errorCode > 500:
				logger.Log.Error().
					Interface("request", logRequest).
					Interface("response", logResponse).
					Msg(logMessage)
			default:
				logger.Log.Warn().Msg(logMessage)
			}
		}),
	})
}
