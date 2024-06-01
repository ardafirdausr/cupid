package http

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"com.ardafirdausr.cupid/internal/pkg/logger"
	"com.ardafirdausr.cupid/internal/repository/mongo/seed"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/mongo"
)

type app struct {
	config  config
	srv     *httpServer
	mongoDB *mongo.Database
}

func newApp(
	config config,
	srv *httpServer,
	router *httpRouter,
	mongoDB *mongo.Database,
) *app {
	router.setupRouteOnServer(srv.echo)
	return &app{config: config, srv: srv, mongoDB: mongoDB}
}

func (app *app) setupDatabase() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	seed.SetupUserCollection(ctx, app.mongoDB)
	seed.SetupMatchingCollection(ctx, app.mongoDB)
	seed.SetupSubscriptionPlanCollection(ctx, app.mongoDB)
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

	app.setupDatabase()
	app.srv.start()
	defer app.srv.close()
	<-ctx.Done()
}
