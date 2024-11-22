package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var C Config

// Loads config from a file.
func Load(configFile string) (Config, error) {
	var cfg Config

	content, err := os.ReadFile(configFile)
	if err != nil {
		return cfg, fmt.Errorf(
			"Failed to load application config file %q: %v\n",
			configFile,
			err,
		)
	}

	if err = yaml.Unmarshal(content, &cfg); err != nil {
		return cfg, fmt.Errorf("Failed to unmarshal config file %q: %v\n",
			configFile,
			err)
	}

	return cfg, nil
}
