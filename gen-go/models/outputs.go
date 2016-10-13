package models

// DefaultInternalError represents a generic 500 response.
type DefaultInternalError struct {
	Msg string `json:"msg"`
}

// Error returns the internal error that caused the 500.
func (d DefaultInternalError) Error() string {
	return d.Msg
}

// DefaultBadRequest represents a generic 400 response. It used internally by Swagger as the
// response when a request fails the validation defined in the Swagger yml file.
type DefaultBadRequest struct {
	Msg string `json:"msg"`
}

// Error returns the validation error that caused the 400.
func (d DefaultBadRequest) Error() string {
	return d.Msg
}

// HealthCheckError defines the error interface for HealthCheck.
type HealthCheckError interface {
	error // Extend the error interface
	HealthCheckStatusCode() int
}

// LocationForIPError defines the error interface for LocationForIP.
type LocationForIPError interface {
	error // Extend the error interface
	LocationForIPStatusCode() int
}

// LocationForIP404Output defines the 404 status code response for LocationForIP.
type LocationForIP404Output struct{}

// LocationForIPStatusCode returns the status code for the operation.
func (o LocationForIP404Output) LocationForIPStatusCode() int {
	return 404
}

// Error returns "Status Code: X". We implemented in to satisfy the error
// interface. For a more descriptive error message see the output type.
func (o LocationForIP404Output) Error() string {
	return "Status Code: 404"
}
