package http

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/hantdev/athena"
	"github.com/hantdev/athena/auth"
	"github.com/hantdev/athena/auth/api/http/keys"
	"github.com/hantdev/athena/auth/api/http/pats"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// MakeHandler returns a HTTP handler for API endpoints.
func MakeHandler(svc auth.Service, logger *slog.Logger, instanceID string) http.Handler {
	mux := chi.NewRouter()

	mux = keys.MakeHandler(svc, mux, logger)
	mux = pats.MakeHandler(svc, mux, logger)

	mux.Get("/health", athena.Health("auth", instanceID))
	mux.Handle("/metrics", promhttp.Handler())

	return mux
}
