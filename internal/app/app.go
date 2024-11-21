package app

import (
	"app/internal/config"
)

type App struct {
	config *config.Config
}

func New(config *config.Config) *App {
	app := &App{config: config}

	return app
}
