package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/4lexvav/hit-the-goal/handlers/middlewares"
	"github.com/4lexvav/hit-the-goal/handlers/projects"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middlewares.RequestID)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		binBody, _ := json.Marshal("Hello World")
		w.Write(binBody)
	})

	r.Route("/projects", func(r chi.Router) {
		r.Get("/", projects.GetList)
		r.Post("/", projects.Create)
	})

	return r
}
