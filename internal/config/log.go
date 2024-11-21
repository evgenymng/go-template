package config

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type EncoderConfig struct {
	MessageKey    string               `yaml:"message_key"`
	LevelKey      string               `yaml:"level_key"`
	LevelEncoder  zapcore.LevelEncoder `yaml:"level_encoder"`
	TimeKey       string               `yaml:"time_key"`
	TimeEncoder   zapcore.TimeEncoder  `yaml:"time_encoder"`
	NameKey       string               `yaml:"name_key"`
	CallerKey     string               `yaml:"caller_key"`
	FunctionKey   string               `yaml:"function_key"`
	StacktraceKey string               `yaml:"stacktrace_key"`
}

type LogConfig struct {
	Level            zap.AtomicLevel `yaml:"level"`
	Encoding         string          `yaml:"encoding"`
	OutputPaths      []string        `yaml:"output_paths"`
	ErrorOutputPaths []string        `yaml:"error_output_paths"`
	ShowFileLine     bool            `yaml:"show_file_line"`
	EncoderConfig    EncoderConfig   `yaml:"encoder_config"`
}
