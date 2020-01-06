package main

import (
	"context"
	"flag"
	"log"

	"github.com/Clever/resolve-ip/gen-go/models"
	"github.com/Clever/resolve-ip/gen-go/server"
	"gopkg.in/Clever/kayvee-go.v5/logger"
)

var addr = flag.String("addr", ":8080", "Address to listen at")
var path = flag.String("path", "./GeoLiteCity", "path to Geo Lite City db")

type handler struct {
	db *GeoDB
}

func (h handler) HealthCheck(ctx context.Context) error {
	return nil
}

func (h handler) LocationForIP(ctx context.Context, ip string) (*models.IP, error) {
	latlon, err := h.db.Lookup(ip)
	if err == ErrIPMissing {
		return nil, models.NotFound{Message: "Cannot locate IP"}
	}
	if err == ErrBadIP {
		return nil, models.BadRequest{Message: err.Error()}
	}
	if err != nil {
		return nil, err
	}
	return &models.IP{
		Lat: &latlon.Lat,
		Lon: &latlon.Lon,
	}, nil

}

func main() {
	flag.Parse()
	logger := logger.New("resolve-ip")

	logger.Debug("build-db-start")
	db, err := NewGeoDB(*path)
	if err != nil {
		log.Fatalf("failed to create geo db: %s", err)
	}
	logger.Debug("build-db-end")

	s := server.New(handler{db}, *addr)
	// Serve should not return
	log.Fatal(s.Serve())
}
