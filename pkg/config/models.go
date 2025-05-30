package config

import (
	"time"

	"go.uber.org/zap"
)

type Config struct {
	Release string       `env:"RELEASE, default=unknown_release"`
	Debug   bool         `env:"DEBUG"`
	Log     LogConfig    `env:", prefix=LOG_"`
	Server  ServerConfig `env:", prefix=SERVER_"`
}

type ServerConfig struct {
	Host            string        `env:"HOST, default=0.0.0.0"`
	Port            uint16        `env:"PORT, default=8080"`
	ShutdownTimeout time.Duration `env:"SHUTDOWN_TIMEOUT, default=5s"`
}

type LogConfig struct {
	Level zap.AtomicLevel `env:"LEVEL, default=info"`
}
