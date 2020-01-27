package main

import (
	"email-validator/api"
	"net/http"
	"os"
)

func main() {
	// Starts server
	port := os.Getenv("PORT")
	http.ListenAndServe(":"+port, api.RegisterRoutes())
}
