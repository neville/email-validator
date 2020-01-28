package main

import (
	"email-validator/api"
	"fmt"
	"net/http"
)

func main() {
	// Starts server
	port := "8080" //os.Getenv("PORT")
	fmt.Println("Server listening on port - " + port)

	http.ListenAndServe(":"+port, api.RegisterRoutes())
}
