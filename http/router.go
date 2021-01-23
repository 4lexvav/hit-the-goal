package handlers

import (
	"github.com/4lexvav/hit-the-goal/http/handlers/lists"
	"github.com/4lexvav/hit-the-goal/http/handlers/projects"
	"github.com/4lexvav/hit-the-goal/http/handlers/tasks"
	"github.com/4lexvav/hit-the-goal/http/middlewares"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middlewares.RequestID)

	addProjectRoutes(r)
	addListRoutes(r)
	addTaskRoutes(r)

	return r
}

func addProjectRoutes(r chi.Router) {
	r.Route("/projects", func(r chi.Router) {
		r.Get("/", projects.GetList)
		r.Post("/", projects.Create)

		r.Route("/{projectID}", func(r chi.Router) {
			r.Get("/", projects.GetById)
			r.Patch("/", projects.Update)
			r.Delete("/", projects.Delete)
		})
	})
}

func addListRoutes(r chi.Router) {
	r.Route("/projects/{{projectID}}/lists", func(r chi.Router) {
		r.Get("/", lists.GetList)
		r.Post("/", lists.Create)

		r.Route("/{listID}", func(r chi.Router) {
			r.Get("/", lists.GetById)
			r.Patch("/", lists.Update)
			r.Delete("/", lists.Delete)
		})
	})
}

func addTaskRoutes(r chi.Router) {
	r.Route("/lists/{listID}/tasks", func(r chi.Router) {
		r.Get("/", tasks.GetList)
		r.Post("/", tasks.Create)

		r.Route("/{taskID}", func(r chi.Router) {
			r.Get("/", tasks.GetById)
			r.Patch("/", tasks.Update)
			r.Delete("/", tasks.Delete)
		})
	})
}
