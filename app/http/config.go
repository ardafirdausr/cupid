package http

import (
	"time"

	"com.ardafirdausr.cupid/internal/entity"
	"com.ardafirdausr.cupid/internal/pkg/mongo"
	"github.com/spf13/viper"
)

type config struct {
	common entity.CommonConfig
	http   httpConfig
	mongo  mongo.Config
}

func loadConfig() config {
	viper.SetConfigFile("./app/http/.env")

	viper.ReadInConfig()

	// Override with os env
	viper.AutomaticEnv()
	return config{
		common: loadCommonConfig(),
		http:   loadHTTPConfig(),
		mongo:  loadMongoConfig(),
	}
}

func loadCommonConfig() entity.CommonConfig {
	appName := viper.GetString("APP_NAME")
	if appName == "" {
		appName = "cupid"
	}

	env := viper.GetString("ENV")
	if env == "" {
		env = "production"
	}

	return entity.CommonConfig{AppName: appName}
}

func loadHTTPConfig() httpConfig {
	port := viper.GetInt("HTTP_PORT")
	if port <= 0 {
		port = 8080
	}

	timeout := viper.GetInt("HTTP_TIMEOUT")
	if timeout <= 0 {
		timeout = 30000
	}

	return httpConfig{port: port, timeout: time.Duration(timeout) * time.Millisecond}
}
func loadMongoConfig() mongo.Config {
	uri := viper.GetString("MONGO_URI")
	if uri == "" {
		uri = "mongodb://localhost:27017/cupid"
	}

	database := viper.GetString("MONGO_DATABASE")
	if database == "" {
		database = "cupid"
	}

	timeout := viper.GetInt("MONGO_TIMEOUT")
	if timeout <= 0 {
		timeout = 30000
	}

	minPool := viper.GetInt("MONGO_POOL_MIN_SIZE")
	if minPool <= 0 {
		minPool = 10
	}

	maxPool := viper.GetInt("MONGO_POOL_MAX_SIZE")
	if maxPool <= 0 {
		maxPool = 50
	}

	maxIdleTimePool := viper.GetInt("MONGO_POOL_IDLE_TIMEOUT")
	if maxIdleTimePool <= 0 {
		maxIdleTimePool = 30000
	}

	return mongo.Config{
		URI:             uri,
		DB:              database,
		MinPool:         uint64(minPool),
		MaxPool:         uint64(maxPool),
		MaxIdleTimePool: time.Duration(maxIdleTimePool) * time.Millisecond,
	}
}
