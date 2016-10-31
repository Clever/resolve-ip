package client

import (
	"context"

	"github.com/Clever/resolve-ip/gen-go/models"
)

//go:generate $GOPATH/bin/mockgen -source=$GOFILE -destination=mock_client.go -package=client

// Client defines the methods available to clients of the resolve-ip service.
type Client interface {

	// HealthCheck makes a GET request to /healthcheck.
	// Checks if the service is healthy
	HealthCheck(ctx context.Context) error

	// LocationForIP makes a GET request to /ip/{ip}.
	// Gets the lat/lon for a given IP.
	LocationForIP(ctx context.Context, i *models.LocationForIPInput) (*models.IP, error)
}
