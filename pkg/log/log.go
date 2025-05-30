package log

import (
	"go-template/pkg/config"

	"go.uber.org/zap"
)

var S *zap.SugaredLogger

// Creates a new logger instance from the configuration.
func NewLogger(cfg config.Config) *zap.SugaredLogger {
	var conf zap.Config

	if cfg.Debug {
		conf = zap.NewDevelopmentConfig()
		conf.Encoding = "console"
	} else {
		conf = zap.NewProductionConfig()
	}
	conf.Level = cfg.Log.Level

	return zap.Must(conf.Build()).Sugar()
}
