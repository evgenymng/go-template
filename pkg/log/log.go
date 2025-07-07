package log

import (
	"go-template/pkg/config"

	"go.uber.org/zap"
)

// Creates a new logger instance from the configuration and replaces the global logger.
func Configure(cfg config.Config) {
	var conf zap.Config

	if cfg.Dev {
		conf = zap.NewDevelopmentConfig()
		conf.Encoding = "console"
	} else {
		conf = zap.NewProductionConfig()
	}
	conf.Level = cfg.Log.Level

	logger := zap.Must(conf.Build())
	zap.ReplaceGlobals(logger)
}
