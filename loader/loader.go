package loader

import (
	"log"

	"github.com/4lexvav/hit-the-goal/config"
	"github.com/4lexvav/hit-the-goal/logger"
	"github.com/4lexvav/hit-the-goal/services"
	"github.com/4lexvav/hit-the-goal/store/repo"
	"github.com/4lexvav/hit-the-goal/store/repo/postgres"
	"github.com/4lexvav/hit-the-goal/validator"
)

const defaultConfigPath = "config.json"

func Load() {
	if err := config.Load(defaultConfigPath); err != nil {
		log.Fatalf("Failed to load config %s", err.Error())
	}

	if err := logger.Load(config.Get().LogPreset); err != nil {
		log.Fatalf("Failed to load logger: %s", err.Error())
	}

	if err := validator.Load(); err != nil {
		logger.Get().Fatalw("Failed to load validator", "error", err)
	}

	if err := postgres.Load(config.Get().Postgres, logger.Get()); err != nil {
		logger.Get().Fatalw("Failed to connect to postgres", "error", err)
	}

	if err := repo.Load(); err != nil {
		logger.Get().Fatalw("Failed to initialize postgres repository", "error", err)
	}

	if err := services.Load(); err != nil {
		logger.Get().Fatalw("Failed to load services", "error", err)
	}
}
