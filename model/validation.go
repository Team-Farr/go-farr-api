package model

import "net/http"

//ApiValidation - interface for all submissions to the API used for deserialization and validation, while still maintaining our context through the http.Request
type ApiValidation interface {
	Validate(r *http.Request) error
}
