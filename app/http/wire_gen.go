// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package http

import (
	"com.ardafirdausr.cupid/app/http/handler"
	"com.ardafirdausr.cupid/internal"
	"com.ardafirdausr.cupid/internal/pkg/helper"
	"com.ardafirdausr.cupid/internal/pkg/mongo"
	"com.ardafirdausr.cupid/internal/pkg/validator"
	mongo2 "com.ardafirdausr.cupid/internal/repository/mongo"
	"com.ardafirdausr.cupid/internal/service"
	"github.com/google/wire"
)

// Injectors from injector.go:

func InitializeApp() (*app, func(), error) {
	config2 := _wireConfigValue
	httpHttpConfig := config2.http
	httpHttpServer, err := newHTTPServer(httpHttpConfig)
	if err != nil {
		return nil, nil, err
	}
	commonConfig := config2.common
	mongoConfig := config2.mongo
	database, cleanup, err := mongo.NewMongoDatabase(mongoConfig)
	if err != nil {
		return nil, nil, err
	}
	userMongoRepository := mongo2.NewUserMongoRepository(database)
	userService := service.NewUserService(userMongoRepository)
	goPlaygroundValidator := validator.NewGoPlayValidator()
	userHandler := handler.NewUserHandler(userService, goPlaygroundValidator)
	authService := service.NewAuthService(commonConfig, userMongoRepository)
	authHandler := handler.NewAuthHandler(authService, goPlaygroundValidator)
	matchingMongoRepositry := mongo2.NewMatchingMongoRepository(database)
	matchingService := service.NewMatchingService(matchingMongoRepositry)
	injector := helper.Newinjector(userService)
	matchingHandler := handler.NewMatchingHandler(matchingService, injector)
	httpHttpRouter := newRouter(commonConfig, userHandler, authHandler, matchingHandler)
	httpApp := newApp(config2, httpHttpServer, httpHttpRouter)
	return httpApp, func() {
		cleanup()
	}, nil
}

var (
	_wireConfigValue = cfg
)

// injector.go:

var cfg = loadConfig()

var configSet = wire.NewSet(wire.Value(cfg), wire.FieldsOf(
	new(config),
	"common",
	"http",
	"mongo"),
)

var handlerSet = wire.NewSet(handler.NewUserHandler, handler.NewAuthHandler, handler.NewMatchingHandler)

var serviceSet = wire.NewSet(service.NewUserService, wire.Bind(new(internal.UserServicer), new(*service.UserService)), service.NewAuthService, wire.Bind(new(internal.AuthServicer), new(*service.AuthService)), service.NewMatchingService, wire.Bind(new(internal.MatchingServicer), new(*service.MatchingService)))

var repoSet = wire.NewSet(mongo2.NewUserMongoRepository, wire.Bind(new(internal.UserRepositorier), new(*mongo2.UserMongoRepository)), mongo2.NewMatchingMongoRepository, wire.Bind(new(internal.MatchingRepositorier), new(*mongo2.MatchingMongoRepositry)))

var driverSet = wire.NewSet(mongo.NewMongoDatabase)

var pkgSet = wire.NewSet(validator.NewGoPlayValidator, wire.Bind(new(validator.Validator), new(*validator.GoPlaygroundValidator)), helper.Newinjector)
