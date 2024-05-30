package http

import (
	"context"
	"fmt"
	"time"

	"com.ardafirdausr.cupid/app/http/handler"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

type httpServer struct {
	echo   *echo.Echo
	logger *zerolog.Logger
	port   int
}

func newHTTPServer(port int, logger *zerolog.Logger) *httpServer {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.HTTPErrorHandler = handler.ErrorHandler

	return &httpServer{
		echo:   echo.New(),
		logger: logger,
		port:   port,
	}
}

func (srv *httpServer) start() {
	srv.logger.Info().Msg(fmt.Sprintf("http: starting server on port %d", srv.port))
	go func() {
		if err := srv.echo.Start(fmt.Sprintf("0.0.0.0:%d", srv.port)); err != nil {
			srv.logger.Info().Msg(err.Error())
		}
	}()
}

func (srv *httpServer) close() {
	ctxShutdown, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	srv.logger.Info().Msg("Closing http server")
	if err := srv.echo.Shutdown(ctxShutdown); err != nil {
		srv.logger.Info().Msg(fmt.Sprintf("Failed to close http server: %v", err))
	}
}
