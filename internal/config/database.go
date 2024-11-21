package config

type DatabaseConfig struct {
	Username string `yaml:"username"`
	Name     string `yaml:"name"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
}
