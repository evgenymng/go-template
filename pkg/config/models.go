package config

import (
	"go.uber.org/zap"
)

type Config struct {
	Version string       `env:"RELEASE, default=unknown_release"`
	Dev     bool         `env:"DEV"`
	Log     LogConfig    `env:", prefix=LOG_"`
	Server  ServerConfig `env:", prefix=SERVER_"`
	Otel    OtelConfig   `env:", prefix=OTEL_"`
}

type ServerConfig struct {
	Host string `env:"HOST, default=0.0.0.0"`
	Port uint16 `env:"PORT, default=8080"`
}

type LogConfig struct {
	Level zap.AtomicLevel `env:"LEVEL, default=info"`
}

type OtelConfig struct {
	Host        string `env:"HOST, default=localhost"`
	Port        uint16 `env:"PORT, default=4317"`
	Secure      bool   `env:"SECURE, default=true"`
	ServiceName string `env:"SERVICE_NAME, default=github.com/evgenymng/go-template"`
}
