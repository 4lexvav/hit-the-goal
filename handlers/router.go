package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		binBody, _ := json.Marshal("Hello World")

		w.Write(binBody)
	})

	return r
}
