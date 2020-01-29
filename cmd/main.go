package main

import (
	"email-validator/api"
	"fmt"
	"net/http"
	"os"
)

func main() {
	// Starts server
	port := os.Getenv("PORT")
	fmt.Println("Server listening on port - " + port)

	http.ListenAndServe(":"+port, api.RegisterRoutes())
}
