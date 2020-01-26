package controller

import "net/http"

// Validate ...
func Validate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Validating email..."))
}
