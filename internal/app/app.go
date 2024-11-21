package app

import (
	"app/internal/config"

	"github.com/redis/go-redis/v9"
)

type App struct {
	config      config.Config
	RedisClient *redis.Client
}

func New(config config.Config) App {
	app := App{config: config}

	return app
}
