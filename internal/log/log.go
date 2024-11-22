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

// Creates a new [Logger] using the global configuration.
func New() *Logger {
	encConf := zapcore.EncoderConfig{
		MessageKey:    config.C.Log.EncoderConfig.MessageKey,
		LevelKey:      config.C.Log.EncoderConfig.LevelKey,
		TimeKey:       config.C.Log.EncoderConfig.TimeKey,
		NameKey:       config.C.Log.EncoderConfig.NameKey,
		CallerKey:     config.C.Log.EncoderConfig.CallerKey,
		FunctionKey:   config.C.Log.EncoderConfig.FunctionKey,
		StacktraceKey: config.C.Log.EncoderConfig.StacktraceKey,
		EncodeLevel:   config.C.Log.EncoderConfig.LevelEncoder,
		EncodeTime:    config.C.Log.EncoderConfig.TimeEncoder,
	}

	conf := zap.Config{
		Level:            config.C.Log.Level,
		Encoding:         config.C.Log.Encoding,
		OutputPaths:      config.C.Log.OutputPaths,
		ErrorOutputPaths: config.C.Log.ErrorOutputPaths,
		Development:      config.C.Log.ShowFileLine,
		EncoderConfig:    encConf,
	}

	return &Logger{
		internal: zap.Must(conf.Build()).Sugar(),
	}
}
