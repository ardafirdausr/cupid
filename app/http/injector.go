//go:build wireinject

package http

import (
	"com.ardafirdausr.cupid/app/http/handler"
	"com.ardafirdausr.cupid/internal"
	customJwt "com.ardafirdausr.cupid/internal/pkg/jwt"
	"com.ardafirdausr.cupid/internal/pkg/mongo"
	"com.ardafirdausr.cupid/internal/pkg/validator"
	mongoRepository "com.ardafirdausr.cupid/internal/repository/mongo"
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
	handler.NewMatchingHandler,
	handler.NewSubscriptionHandler,
)

var serviceSet = wire.NewSet(
	service.NewUserService,
	wire.Bind(new(internal.UserServicer), new(*service.UserService)),
	service.NewAuthService,
	wire.Bind(new(internal.AuthServicer), new(*service.AuthService)),
	service.NewMatchingService,
	wire.Bind(new(internal.MatchingServicer), new(*service.MatchingService)),
	service.NewSubscriptionService,
	wire.Bind(new(internal.SubscriptionServicer), new(*service.SubscriptionServicer)),
)

var repoSet = wire.NewSet(
	mongoRepository.NewUserRepository,
	wire.Bind(new(internal.UserRepositorier), new(*mongoRepository.UserRepository)),
	mongoRepository.NewMatchingRepository,
	wire.Bind(new(internal.MatchingRepositorier), new(*mongoRepository.MatchingRepository)),
	mongoRepository.NewSubscriptionRepository,
	wire.Bind(new(internal.SubscriptionRepositorier), new(*mongoRepository.SubscriptionRepository)),
)

var driverSet = wire.NewSet(
	mongo.NewMongoDatabase,
)

var pkgSet = wire.NewSet(
	validator.NewGoPlayValidator,
	wire.Bind(new(validator.Validator), new(*validator.GoPlaygroundValidator)),
	customJwt.NewHelper,
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
