//go:build wireinject

package http

import (
	"com.ardafirdausr.cupid/app/http/handler"
	"com.ardafirdausr.cupid/internal"
	"com.ardafirdausr.cupid/internal/pkg/mongo"
	"com.ardafirdausr.cupid/internal/pkg/validator"
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
	handler.NewAuthHandler,
)

var serviceSet = wire.NewSet(
	service.NewUserService,
	wire.Bind(new(internal.UserServicer), new(*service.UserService)),
	service.NewAuthService,
	wire.Bind(new(internal.AuthServicer), new(*service.AuthService)),
)

var repoSet = wire.NewSet(
	repository.NewUserMongoRepository,
	wire.Bind(new(internal.UserRepositorier), new(*repository.UserMongoRepository)),
)

var driverSet = wire.NewSet(
	mongo.NewMongoDatabase,
)

var pkgSet = wire.NewSet(
	validator.NewGoPlayValidator,
	wire.Bind(new(validator.Validator), new(*validator.GoPlaygroundValidator)),
)

func InitializeApp() (*app, func(), error) {
	wire.Build(
		configSet,
		handlerSet,
		repoSet,
		serviceSet,
		driverSet,
		pkgSet,
		newHTTPServer,
		newRouter,
		newApp)
	return nil, nil, nil
}
