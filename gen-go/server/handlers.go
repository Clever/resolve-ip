package server

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/Clever/kayvee-go/v7/logger"
	"github.com/Clever/resolve-ip/gen-go/models/v4"
	"github.com/go-errors/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/gorilla/mux"
	"golang.org/x/xerrors"
)

var _ = strconv.ParseInt
var _ = strfmt.Default
var _ = swag.ConvertInt32
var _ = errors.New
var _ = mux.Vars
var _ = bytes.Compare
var _ = ioutil.ReadAll

var formats = strfmt.Default
var _ = formats

// convertBase64 takes in a string and returns a strfmt.Base64 if the input
// is valid base64 and an error otherwise.
func convertBase64(input string) (strfmt.Base64, error) {
	temp, err := formats.Parse("byte", input)
	if err != nil {
		return strfmt.Base64{}, err
	}
	return *temp.(*strfmt.Base64), nil
}

// convertDateTime takes in a string and returns a strfmt.DateTime if the input
// is a valid DateTime and an error otherwise.
func convertDateTime(input string) (strfmt.DateTime, error) {
	temp, err := formats.Parse("date-time", input)
	if err != nil {
		return strfmt.DateTime{}, err
	}
	return *temp.(*strfmt.DateTime), nil
}

// convertDate takes in a string and returns a strfmt.Date if the input
// is a valid Date and an error otherwise.
func convertDate(input string) (strfmt.Date, error) {
	temp, err := formats.Parse("date", input)
	if err != nil {
		return strfmt.Date{}, err
	}
	return *temp.(*strfmt.Date), nil
}

func jsonMarshalNoError(i interface{}) string {
	bytes, err := json.Marshal(i)
	if err != nil {
		// This should never happen
		return ""
	}
	return string(bytes)
}

// statusCodeForHealthCheck returns the status code corresponding to the returned
// object. It returns -1 if the type doesn't correspond to anything.
func statusCodeForHealthCheck(obj interface{}) int {

	switch obj.(type) {

	case *models.BadRequest:
		return 400

	case *models.InternalError:
		return 500

	case models.BadRequest:
		return 400

	case models.InternalError:
		return 500

	default:
		return -1
	}
}

func (h handler) HealthCheckHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	err := h.HealthCheck(ctx)

	if err != nil {
		logger.FromContext(ctx).AddContext("error", err.Error())
		if btErr, ok := err.(*errors.Error); ok {
			logger.FromContext(ctx).AddContext("stacktrace", string(btErr.Stack()))
		} else if xerr, ok := err.(xerrors.Formatter); ok {
			logger.FromContext(ctx).AddContext("frames", fmt.Sprintf("%+v", xerr))
		}
		statusCode := statusCodeForHealthCheck(err)
		if statusCode == -1 {
			err = models.InternalError{Message: err.Error()}
			statusCode = 500
		}
		http.Error(w, jsonMarshalNoError(err), statusCode)
		return
	}

	w.WriteHeader(200)
	w.Write([]byte(""))

}

// newHealthCheckInput takes in an http.Request an returns the input struct.
func newHealthCheckInput(r *http.Request) (*models.HealthCheckInput, error) {
	var input models.HealthCheckInput

	var err error
	_ = err

	return &input, nil
}

// statusCodeForLocationForIP returns the status code corresponding to the returned
// object. It returns -1 if the type doesn't correspond to anything.
func statusCodeForLocationForIP(obj interface{}) int {

	switch obj.(type) {

	case *models.BadRequest:
		return 400

	case *models.IP:
		return 200

	case *models.InternalError:
		return 500

	case *models.NotFound:
		return 404

	case models.BadRequest:
		return 400

	case models.IP:
		return 200

	case models.InternalError:
		return 500

	case models.NotFound:
		return 404

	default:
		return -1
	}
}

func (h handler) LocationForIPHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	ip, err := newLocationForIPInput(r)
	if err != nil {
		logger.FromContext(ctx).AddContext("error", err.Error())
		http.Error(w, jsonMarshalNoError(models.BadRequest{Message: err.Error()}), http.StatusBadRequest)
		return
	}

	err = models.ValidateLocationForIPInput(ip)

	if err != nil {
		logger.FromContext(ctx).AddContext("error", err.Error())
		http.Error(w, jsonMarshalNoError(models.BadRequest{Message: err.Error()}), http.StatusBadRequest)
		return
	}

	resp, err := h.LocationForIP(ctx, ip)

	if err != nil {
		logger.FromContext(ctx).AddContext("error", err.Error())
		if btErr, ok := err.(*errors.Error); ok {
			logger.FromContext(ctx).AddContext("stacktrace", string(btErr.Stack()))
		} else if xerr, ok := err.(xerrors.Formatter); ok {
			logger.FromContext(ctx).AddContext("frames", fmt.Sprintf("%+v", xerr))
		}
		statusCode := statusCodeForLocationForIP(err)
		if statusCode == -1 {
			err = models.InternalError{Message: err.Error()}
			statusCode = 500
		}
		http.Error(w, jsonMarshalNoError(err), statusCode)
		return
	}

	respBytes, err := json.Marshal(resp)
	if err != nil {
		logger.FromContext(ctx).AddContext("error", err.Error())
		http.Error(w, jsonMarshalNoError(models.InternalError{Message: err.Error()}), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCodeForLocationForIP(resp))
	w.Write(respBytes)

}

// newLocationForIPInput takes in an http.Request an returns the ip parameter
// that it contains. It returns an error if the request doesn't contain the parameter.
func newLocationForIPInput(r *http.Request) (string, error) {
	ip := mux.Vars(r)["ip"]
	if len(ip) == 0 {
		return "", errors.New("Parameter ip must be specified")
	}
	return ip, nil
}
