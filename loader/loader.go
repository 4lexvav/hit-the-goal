package loader

import (
	"log"

	"github.com/4lexvav/hit-the-goal/config"
)

const defaultConfigPath = "config.json"

func Load() {
	if err := config.Load(defaultConfigPath); err != nil {
		log.Fatalf("Failed to load config %s", err.Error())
	}
}
