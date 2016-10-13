package server

import (
	"context"

	"github.com/Clever/resolve-ip/gen-go/models"
)

//go:generate $GOPATH/bin/mockgen -source=$GOFILE -destination=mock_controller.go -package=server

// Controller defines the interface for the resolve-ip service.
type Controller interface {

	// HealthCheck makes a GET request to /healthcheck.
	// Checks if the service is healthy
	HealthCheck(ctx context.Context) error

	// LocationForIP makes a GET request to /ip/{ip}.
	// Gets the lat/lon for a given IP.
	LocationForIP(ctx context.Context, i *models.LocationForIPInput) (*models.IP, error)
}
