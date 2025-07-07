package log

import (
	"go-template/pkg/config"

	"github.com/uptrace/opentelemetry-go-extra/otelzap"
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

	// wrap with OpenTelemetry integration
	otelLogger := otelzap.New(logger)
	otelzap.ReplaceGlobals(otelLogger)
}
