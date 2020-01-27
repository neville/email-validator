package controller

import (
	"email-validator/api/module"
	"encoding/json"
	"net/http"
	"strings"

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
	Reason string `json:"reason"`
}
type smtpValidator struct {
	Valid  bool   `json:"valid"`
	Reason string `json:"reason"`
}

// Validate ...
func Validate(w http.ResponseWriter, r *http.Request) {
	req, err := bindRequest(r)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		return
	}

	response := &Response{}
	isValid := false

	// Regex validation
	isFormatValid, err := module.ValidateFormat(req.Email)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		return
	}
	response.Validators.Regexp.Valid = isFormatValid

	// Extracts domain
	addressSymbolIndex := strings.Index(req.Email, "@")
	if addressSymbolIndex == -1 {
		render.Status(r, http.StatusBadRequest)
		return
	}
	domain := req.Email[addressSymbolIndex+1:]

	// Domain validation
	isDomainValid, err := module.ValidateDomain(domain)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		return
	}
	response.Validators.Domain.Valid = isDomainValid

	// SMTP validation
	isSMTPValid, err := module.ValidateSMTP(domain)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		return
	}
	response.Validators.SMTP.Valid = isSMTPValid

	// Overall validation
	response.Valid = isValid

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
