package server

import (
	"context"

	"github.com/Clever/resolve-ip/gen-go/models/v4"
)

//go:generate mockgen -source=$GOFILE -destination=mock_controller.go -package server --build_flags=--mod=mod -imports=models=github.com/Clever/resolve-ip/gen-go/models/v4

// Controller defines the interface for the resolve-ip service.
type Controller interface {

	// HealthCheck handles GET requests to /healthcheck
	// Checks if the service is healthy
	// 200: nil
	// 400: *models.BadRequest
	// 500: *models.InternalError
	// default: client side HTTP errors, for example: context.DeadlineExceeded.
	HealthCheck(ctx context.Context) error

	// LocationForIP handles GET requests to /ip/{ip}
	// Gets the lat/lon for a given IP.
	// 200: *models.IP
	// 400: *models.BadRequest
	// 404: *models.NotFound
	// 500: *models.InternalError
	// default: client side HTTP errors, for example: context.DeadlineExceeded.
	LocationForIP(ctx context.Context, ip string) (*models.IP, error)
}
