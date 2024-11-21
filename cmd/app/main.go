package main

import (
	"fmt"
	"os"

	"app/internal/config"
	"app/internal/log"
)

func main() {
	configFile, ok := os.LookupEnv("CONFIG_FILE")
	if !ok {
		panic(
			"Please, point to the config file in the CONFIG_FILE env variable",
		)
	}

	cfg, err := config.Load(configFile)
	if err != nil {
		panic(fmt.Sprintf("Failed to load the config: %v", err))
	}

	log.S = log.New(cfg.Log)
	log.S.Info(
		"Config is loaded, logger is initialized",
		log.L().Tag(log.TagStartup),
	)
}
