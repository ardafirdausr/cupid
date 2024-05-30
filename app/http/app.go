package http

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/rs/zerolog"
)

const port = 8000

type app struct {
	logger *zerolog.Logger
	srv    *httpServer
}

func InitializeApp() *app {
	zlog := zerolog.New(os.Stdout).With().Timestamp().Logger()
	logger := &zlog
	logger.Info().Msg("Initializing app")

	srv := newHTTPServer(port, logger)

	router := newRouter()
	router.setupRouteOnServer(srv.echo)

	return &app{logger: logger, srv: srv}
}

func (app *app) Start() {
	runtime.GOMAXPROCS(runtime.NumCPU())

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
		msgStr := fmt.Sprintf("SYSTEM CALL: %+v", in)
		app.logger.Info().Msg(msgStr)
		cancel()
	}()

	app.srv.start()
	defer app.srv.close()
	<-ctx.Done()
}
