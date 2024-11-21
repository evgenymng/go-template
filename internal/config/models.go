package config

type Config struct {
	ApiKeys     []string     `yaml:"api_keys"`
	EnableDocs  bool         `yaml:"enable_docs"`
	EnablePprof bool         `yaml:"enable_pprof"`
	Log         LogConfig    `yaml:"log"`
	Server      ServerConfig `yaml:"server"`
	// ClickHouse, Mongo, other integrations...
}

type ServerConfig struct {
	Mode            string `yaml:"mode"` // "debug" or "release"
	Host            string `yaml:"host"`
	Port            uint16 `yaml:"port"`
	ShutdownTimeout uint   `yaml:"shutdown_timeout"`
}
