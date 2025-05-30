package config

import (
	"context"
	"encoding/json"
	"log"

	"github.com/sethvargo/go-envconfig"
)

var C Config

// Loads config from the environment.
func Load(i any) {
	if err := envconfig.Process(context.Background(), i); err != nil {
		log.Fatal(err)
	}
	text, _ := json.MarshalIndent(i, "", "\t")
	log.Printf("Loaded config:\n%s\n\n", string(text))
}
