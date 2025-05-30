package main

import (
	"go-template/internal"
	"go-template/pkg/config"
	"go-template/pkg/log"

	"github.com/joho/godotenv"
)

// @title Backend Service Template
func main() {
	_ = godotenv.Load()

	var cfg config.Config
	config.Load(&cfg)
	config.C = cfg

	log.S = log.NewLogger(config.C)
	log.S.Info("Config is loaded, logger is initialized")

	internal.Launch()
}
