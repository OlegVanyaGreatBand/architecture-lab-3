//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/OlegVanyaGreatBand/architecture-lab-3/server/models"
)

// ComposeApiServer will create an instance of TelemetryApiServer according to providers defined in this file.
func ComposeApiServer(port HttpPortNumber) (*TelemetryApiServer, error) {
	wire.Build(
		// DB connection provider (defined in main.go).
		NewDbConnection,
		// Add providers from models package.
		models.Providers,
		// Provide TelemetryApiServer instantiating the structure and injecting channels handler and port number.
		wire.Struct(new(TelemetryApiServer), "Port", "TelemetryHandler"),
	)
	return nil, nil
}