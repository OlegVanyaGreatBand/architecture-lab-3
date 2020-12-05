
package main

import (
	"context"
	"fmt"
	"github.com/OlegVanyaGreatBand/architecture-lab-3/server/models"
	"net/http"
)

type HttpPortNumber int

// TelemetryApiServer configures necessary handlers and starts listening on a configured port.
type TelemetryApiServer struct {
	Port HttpPortNumber

	TelemetryHandler models.TelemetryHttpHandler

	server *http.Server
}

// Start will set all handlers and start listening.
// If this methods succeeds, it does not return until server is shut down.
// Returned error will never be nil.
func (s *TelemetryApiServer) Start() error {
	if s.TelemetryHandler == nil {
		return fmt.Errorf("channels HTTP handler is not defined - cannot start")
	}
	if s.Port == 0 {
		return fmt.Errorf("port is not defined")
	}

	handler := new(http.ServeMux)
	handler.HandleFunc("/", s.TelemetryHandler)

	s.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.Port),
		Handler: handler,
	}

	return s.server.ListenAndServe()
}

// Stops will shut down previously started HTTP server.
func (s *TelemetryApiServer) Stop() error {
	if s.server == nil {
		return fmt.Errorf("server was not started")
	}
	return s.server.Shutdown(context.Background())
}
