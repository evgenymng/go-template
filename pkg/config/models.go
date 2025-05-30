package config

import (
	"go.uber.org/zap"
)

type Config struct {
	Version string       `env:"RELEASE, default=unknown_release"`
	Dev     bool         `env:"DEV"`
	Log     LogConfig    `env:", prefix=LOG_"`
	Server  ServerConfig `env:", prefix=SERVER_"`
}

type ServerConfig struct {
	Host string `env:"HOST, default=0.0.0.0"`
	Port uint16 `env:"PORT, default=8080"`
}

type LogConfig struct {
	Level zap.AtomicLevel `env:"LEVEL, default=info"`
}
