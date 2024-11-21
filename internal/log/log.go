package log

import (
	"app/internal/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var S *Logger

type Logger struct {
	internal *zap.SugaredLogger
}

func (l *Logger) GetInternal() *zap.SugaredLogger {
	return l.internal
}

// Creates a new [Logger] using the global configuration and level.
func New(cfg config.LogConfig) *Logger {
	encConf := zapcore.EncoderConfig{
		MessageKey:    cfg.EncoderConfig.MessageKey,
		LevelKey:      cfg.EncoderConfig.LevelKey,
		TimeKey:       cfg.EncoderConfig.TimeKey,
		NameKey:       cfg.EncoderConfig.NameKey,
		CallerKey:     cfg.EncoderConfig.CallerKey,
		FunctionKey:   cfg.EncoderConfig.FunctionKey,
		StacktraceKey: cfg.EncoderConfig.StacktraceKey,
		EncodeLevel:   cfg.EncoderConfig.LevelEncoder,
		EncodeTime:    cfg.EncoderConfig.TimeEncoder,
	}

	conf := zap.Config{
		Level:            cfg.Level,
		Encoding:         cfg.Encoding,
		OutputPaths:      cfg.OutputPaths,
		ErrorOutputPaths: cfg.ErrorOutputPaths,
		Development:      cfg.ShowFileLine,
		EncoderConfig:    encConf,
	}

	return &Logger{
		internal: zap.Must(conf.Build()).Sugar(),
	}
}
