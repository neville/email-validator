package controller

import (
	"net/http"

	"github.com/go-chi/render"
)

// Request ...
type Request struct {
	Email string `json:"email"`
}

// Response ...
type Response struct {
	Valid      string `json:"valid"`
	Validators struct {
		Regexp struct {
			Valid bool
		} `json:"regexp"`
		Domain struct {
			Valid  bool
			Reason string
		} `json:"domain"`
		SMTP struct {
			Valid  bool
			Reason string
		} `json:"smtp"`
	} `json:"validators"`
}

// Validate ...
func Validate(w http.ResponseWriter, r *http.Request) {
	// Binds request data
	request := &Request{}
	render.DecodeJSON(r.Body, &request)

	if request.Email == "" {
		render.Status(r, http.StatusNotFound)
	}

	w.Write([]byte("Validating email..."))
}
