//go:build wireinject

package http

import (
	"com.ardafirdausr.cupid/app/http/handler"
	"com.ardafirdausr.cupid/internal"
	"com.ardafirdausr.cupid/internal/pkg/mongo"
	repository "com.ardafirdausr.cupid/internal/repository/mongo"
	"com.ardafirdausr.cupid/internal/service"
	"github.com/google/wire"
)

var cfg = loadConfig()

var configSet = wire.NewSet(
	wire.Value(cfg),
	wire.FieldsOf(
		new(config),
		"common",
		"http",
		"mongo"),
)

var handlerSet = wire.NewSet(
	handler.NewUserHandler,
)

var serviceSet = wire.NewSet(
	wire.Bind(new(internal.UserServicer), new(*service.UserService)),
	service.NewUserService,
)

var repoSet = wire.NewSet(
	wire.Bind(new(internal.UserRepositorier), new(*repository.UserMongoRepository)),
	repository.NewUserMongoRepository,
)

var driverSet = wire.NewSet(
	mongo.NewMongoDatabase,
)

func InitializeApp() (*app, func(), error) {
	wire.Build(
		configSet,
		handlerSet,
		repoSet,
		serviceSet,
		driverSet,
		newHTTPServer,
		newRouter,
		newApp)
	return nil, nil, nil
}
