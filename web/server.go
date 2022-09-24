package web

import (
	"context"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// Config contains the required configration information for the web layer
type Config struct {
	Service string
	Version string
	Host    string
	Port    int
}

// ListenAndServe configures the routes and starts the webserver, along with the gofunc to handle shutdown
func ListenAndServe(ctx context.Context, cfg *Config, service TranslationService) error {
	log.Info("start server")
	r := setupRoutes(cfg, service)
	server := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
	}

	go func() {
		<-ctx.Done()
		if err := server.Shutdown(context.Background()); err != nil {
			log.WithError(err).Info("shutting down")
		}
	}()

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.WithError(err).Info("server failed to listen")
	}
	log.Info("server stopped")
	return nil
}
