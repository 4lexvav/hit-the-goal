package main

import (
	"net/http"

	"github.com/4lexvav/hit-the-goal/config"
	handlers "github.com/4lexvav/hit-the-goal/http"
	"github.com/4lexvav/hit-the-goal/loader"
	"github.com/4lexvav/hit-the-goal/logger"
)

func main() {
	// Load configs
	loader.Load()

	// Listen HTTP requests
	server := &http.Server{
		Addr:    config.Get().ListenURL,
		Handler: handlers.NewRouter(),
	}

	logger.Get().Infow("Listening...", "listening_url", config.Get().ListenURL)
	if err := server.ListenAndServe(); err != nil {
		logger.Get().Errorw("Failed to initialize HTTP server.", "error", err)
	}
}
