package web

import (
	"net/http"

	"github.com/gorilla/mux"
)

const (
	// healthCheckURI is the uri for the basic status endpoint
	healthCheckURI = "/healthcheck"
	linkURI        = "/v1/links"
)

func setupRoutes(cfg *Config, service TranslationService) http.Handler {
	r := mux.NewRouter()
	r.NotFoundHandler = http.NotFoundHandler()

	h := NewLinkHandler(service)
	r.HandleFunc(healthCheckURI, GetHealthCheckHandler(cfg.Service, cfg.Version))

	r.HandleFunc(linkURI, ToHandlerFunc(h.GetLinks)).Methods(http.MethodPost)
	return r
}
