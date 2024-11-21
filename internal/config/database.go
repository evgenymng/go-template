package config

type DatabaseConfig struct {
	Username string `yaml:"username"`
	Database string `yaml:"database"`
	Schema   string `yaml:"schema"` // The default schema. Not used in every database.
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
}
