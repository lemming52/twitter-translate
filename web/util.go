package web

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// ErrorResponse is the struct used to return error messages via the endpoints
type ErrorResponse struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

// EndpointFunc defines the expected signature of any function used as an endpoint
type EndpointFunc func(r *http.Request) (int, interface{}, error)

// ToHandlerFunc converts the endpoint signature to the golang required signature
// and handles writing of content to headers and response bodies.
func ToHandlerFunc(e EndpointFunc) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		code, payload, err := e(r)
		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			w.WriteHeader(code)
			response := errorToResponse(code, err)
			err = json.NewEncoder(w).Encode(response)
			if err != nil {
				log.Fatal("error encoding error")
			}
			return
		}
		w.WriteHeader(code)
		if payload != nil {
			err = json.NewEncoder(w).Encode(payload)
			if err != nil {
				log.Fatal(fmt.Sprintf("error encoding payload: %v", err))
			}
		}
	}
}

// errorToResponse converts a HTTP status code and error description to an ErrorResponse
func errorToResponse(code int, err error) *ErrorResponse {
	return &ErrorResponse{
		Code:        code,
		Description: err.Error(),
	}
}

// HealthCheck is the structure returned by the healthcheck endpoint
type HealthCheck struct {
	Service string `json:"service"`
	Version string `json:"version"`
}

// GetHealthCheckHandler constructs a function to return service health information
func GetHealthCheckHandler(service, version string) func(w http.ResponseWriter, r *http.Request) {
	return ToHandlerFunc(func(r *http.Request) (int, interface{}, error) {
		return http.StatusOK, &HealthCheck{
			Service: service,
			Version: version,
		}, nil
	})
}
