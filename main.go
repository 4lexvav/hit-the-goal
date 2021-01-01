package main

import (
	"net/http"

	"github.com/4lexvav/hit-the-goal/config"
	"github.com/4lexvav/hit-the-goal/handlers"
	"github.com/4lexvav/hit-the-goal/loader"
)

func main() {
	// Load configs
	loader.Load()

	// Listen HTTP requests
	server := &http.Server{
		Addr:    config.Get().ListenURL,
		Handler: handlers.NewRouter(),
	}

	server.ListenAndServe()
}
