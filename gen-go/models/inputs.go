package models

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// These imports may not be used depending on the input parameters
var _ = json.Marshal
var _ = strconv.FormatInt
var _ = validate.Maximum
var _ = strfmt.NewFormats

// HealthCheckInput holds the input parameters for a healthCheck operation.
type HealthCheckInput struct {
}

// Validate returns an error if any of the HealthCheckInput parameters don't satisfy the
// requirements from the swagger yml file.
func (i HealthCheckInput) Validate() error {
	return nil
}

// LocationForIPInput holds the input parameters for a locationForIP operation.
type LocationForIPInput struct {
	Ip string
}

// Validate returns an error if any of the LocationForIPInput parameters don't satisfy the
// requirements from the swagger yml file.
func (i LocationForIPInput) Validate() error {
	return nil
}
