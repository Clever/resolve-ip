package models

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// These imports may not be used depending on the input parameters
var _ = json.Marshal
var _ = fmt.Sprintf
var _ = url.QueryEscape
var _ = strconv.FormatInt
var _ = strings.Replace
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

// Path returns the URI path for the input.
func (i HealthCheckInput) Path() (string, error) {
	path := "/healthcheck"
	urlVals := url.Values{}

	return path + "?" + urlVals.Encode(), nil
}

// LocationForIPInput holds the input parameters for a locationForIP operation.
type LocationForIPInput struct {
	IP string
}

// ValidateLocationForIPInput returns an error if the input parameter doesn't
// satisfy the requirements in the swagger yml file.
func ValidateLocationForIPInput(ip string) error {

	return nil
}

// LocationForIPInputPath returns the URI path for the input.
func LocationForIPInputPath(ip string) (string, error) {
	path := "/ip/{ip}"
	urlVals := url.Values{}

	pathip := ip
	if pathip == "" {
		err := fmt.Errorf("ip cannot be empty because it's a path parameter")
		if err != nil {
			return "", err
		}
	}
	path = strings.Replace(path, "{ip}", pathip, -1)

	return path + "?" + urlVals.Encode(), nil
}
