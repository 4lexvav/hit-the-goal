package handlers

import (
	"github.com/4lexvav/hit-the-goal/http/handlers/projects"
	"github.com/4lexvav/hit-the-goal/http/middlewares"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middlewares.RequestID)

	r.Route("/projects", func(r chi.Router) {
		r.Get("/", projects.GetList)
		r.Post("/", projects.Create)

		r.Route("/{projectID}", func(r chi.Router) {
			r.Get("/", projects.GetById)
			r.Patch("/", projects.Update)
			r.Delete("/", projects.Delete)
		})
	})

	return r
}
