package handlers

import (
	"github.com/4lexvav/hit-the-goal/http/handlers/lists"
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

			AddListRoutes(r)
		})
	})

	return r
}

func AddListRoutes(r chi.Router) {
	r.Route("/lists", func(r chi.Router) {
		r.Get("/", lists.GetList)
		r.Post("/", lists.Create)

		r.Route("/{listID}", func(r chi.Router) {
			r.Get("/", lists.GetById)
			r.Patch("/", lists.Update)
			r.Delete("/", lists.Delete)
		})
	})
}
