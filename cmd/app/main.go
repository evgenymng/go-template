package main

import (
	"fmt"
	"os"

	"app/internal/app"
	"app/internal/config"
	"app/internal/log"
)

//	@title Backend Service
//	@version

// @securityDefinitions.apikey	ApiKeyAuth
// @in							query
// @name						api_key
func main() {
	configFile, ok := os.LookupEnv("CONFIG_FILE")
	if !ok {
		panic(
			"Please, point to the config file in the CONFIG_FILE env variable",
		)
	}

	var err error
	config.C, err = config.Load(configFile)
	if err != nil {
		panic(fmt.Sprintf("Failed to load the config: %v", err))
	}

	log.S = log.New()
	log.S.Info(
		"Config is loaded, logger is initialized",
		log.L(),
	)

	app.Launch()
}
