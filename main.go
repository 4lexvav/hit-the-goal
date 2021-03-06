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
		Addr:    ":" + config.Get().Port,
		Handler: handlers.NewRouter(),
	}

	logger.Get().Infow("Listening on port...", "port", config.Get().Port)
	if err := server.ListenAndServe(); err != nil {
		logger.Get().Errorw("Failed to initialize HTTP server.", "error", err)
	}
}
