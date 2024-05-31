package http

import (
	"context"
	"fmt"
	"time"

	"com.ardafirdausr.cupid/app/http/handler"
	mid "com.ardafirdausr.cupid/app/http/middleware"
	"com.ardafirdausr.cupid/internal/pkg/logger"
	"github.com/labstack/echo/v4"
)

type httpServer struct {
	echo   *echo.Echo
	config httpConfig
}

type httpConfig struct {
	port    int
	timeout time.Duration
}

func newHTTPServer(config httpConfig) (*httpServer, error) {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.HTTPErrorHandler = handler.ErrorHandler
	e.Use(mid.CORSMiddleware())
	e.Use(mid.DumpLogMiddleware())
	e.Use(mid.TimeoutMiddleware(config.timeout))
	e.Use(mid.RecoverMiddleware())
	srv := &httpServer{echo: e, config: config}
	return srv, nil
}

func (srv *httpServer) start() {
	logger.Log.Info().Msg(fmt.Sprintf("http: starting server on port %d", srv.config.port))
	go func() {
		if err := srv.echo.Start(fmt.Sprintf("0.0.0.0:%d", srv.config.port)); err != nil {
			logger.Log.Info().Msg(err.Error())
		}
	}()
}

func (srv *httpServer) close() {
	ctxShutdown, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	logger.Log.Info().Msg("Closing http server")
	if err := srv.echo.Shutdown(ctxShutdown); err != nil {
		logger.Log.Info().Msg(fmt.Sprintf("Failed to close http server: %v", err))
	}
}
