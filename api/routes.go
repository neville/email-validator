package api

import (
	"email-validator/api/controller"
	"net/http"

	"github.com/go-chi/chi"
)

// RegisterRoutes ...
func RegisterRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong..."))
	})

	router.Post("/email/validate", controller.Validate)

	return router
}
