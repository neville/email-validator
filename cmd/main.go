package main

import (
	"email-validation/api/controller"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

func main() {
	port := os.Getenv("PORT")

	// Starts server
	http.ListenAndServe(":"+port, registerRoutes())
}

func registerRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong..."))
	})

	router.Post("/email/validate", controller.Validate)

	return router
}
