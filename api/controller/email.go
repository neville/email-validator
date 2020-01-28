package controller

import (
	"email-validator/api/module"
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"
)

// Request ...
type Request struct {
	Email string `json:"email"`
}

// Response ...
type Response struct {
	Valid      bool            `json:"valid"`
	Validators validationTypes `json:"validators"`
}
type validationTypes struct {
	Regexp regexValidator  `json:"regexp"`
	Domain domainValidator `json:"domain"`
	SMTP   smtpValidator   `json:"smtp"`
}
type regexValidator struct {
	Valid bool `json:"valid"`
}
type domainValidator struct {
	Valid  bool   `json:"valid"`
	Reason string `json:"reason,omitempty"`
}
type smtpValidator struct {
	Valid  bool   `json:"valid"`
	Reason string `json:"reason,omitempty"`
}

// Validate ...
func Validate(w http.ResponseWriter, r *http.Request) {
	req, err := bindRequest(r)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		return
	}

	response := &Response{}

	// Regex validation
	response.Validators.Regexp.Valid, err = module.ValidateFormat(req.Email)
	if err != nil {
		response.Valid = false
		return
	}

	// Domain validation
	response.Validators.Domain.Valid, err = module.ValidateDomain(req.Email)
	if err != nil {
		response.Valid = false
		response.Validators.Domain.Reason = err.Error()
	}

	// SMTP validation
	response.Validators.SMTP.Valid, err = module.ValidateSMTP(req.Email)
	if err != nil {
		response.Valid = false
		response.Validators.SMTP.Reason = err.Error()
	}

	// Overall validation
	if response.Validators.Regexp.Valid &&
		response.Validators.Domain.Valid &&
		response.Validators.SMTP.Valid {
		response.Valid = true
	}

	// Returns
	sendSuccessResponse(w, r, http.StatusOK, response)
}

func bindRequest(r *http.Request) (request *Request, err error) {
	request = &Request{}
	render.DecodeJSON(r.Body, &request)
	if err != nil {
		return nil, err
	}

	return request, nil
}

func sendSuccessResponse(w http.ResponseWriter, r *http.Request, code int, res *Response) {
	bytes, err := json.Marshal(res)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}
