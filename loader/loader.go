package loader

import (
	"log"

	"github.com/4lexvav/hit-the-goal/config"
	"github.com/4lexvav/hit-the-goal/logger"
)

const defaultConfigPath = "config.json"

func Load() {
	if err := config.Load(defaultConfigPath); err != nil {
		log.Fatalf("Failed to load config %s", err.Error())
	}

	if err := logger.Load(config.Get().LogPreset); err != nil {
		log.Fatalf("Failed to load logger: %s", err.Error())
	}
}
