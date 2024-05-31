package http

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"com.ardafirdausr.cupid/internal/pkg/logger"
	"github.com/rs/zerolog"
)

type app struct {
	config config
	srv    *httpServer
}

func newApp(
	config config,
	srv *httpServer,
	router *httpRouter,
) *app {
	router.setupRouteOnServer(srv.echo)
	return &app{config: config, srv: srv}
}

func (app *app) Start() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	logger.SetLogLevel(zerolog.InfoLevel)
	if app.config.common.Environment == "development" {
		logger.SetLogLevel(zerolog.DebugLevel)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	terminalHandler := make(chan os.Signal, 1)
	signal.Notify(
		terminalHandler,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	go func() {
		in := <-terminalHandler
		msgStr := fmt.Sprintf("System Call: %+v", in)
		logger.Log.Info().Msg(msgStr)
		cancel()
	}()

	app.srv.start()
	defer app.srv.close()
	<-ctx.Done()
}
