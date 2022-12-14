// Code generated by goa v3.10.1, DO NOT EDIT.
//
// demo HTTP client types
//
// Command:
// $ goa gen github.com/joberly/demo-go-api/design

package client

import (
	demoviews "github.com/joberly/demo-go-api/gen/demo/views"
	goa "goa.design/goa/v3/pkg"
)

// RandResponseBody is the type of the "demo" service "rand" endpoint HTTP
// response body.
type RandResponseBody struct {
	Result *int64 `form:"result,omitempty" json:"result,omitempty" xml:"result,omitempty"`
}

// RandInvalidArgumentsResponseBody is the type of the "demo" service "rand"
// endpoint HTTP response body for the "invalid_arguments" error.
type RandInvalidArgumentsResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// NewRandResultViewOK builds a "demo" service "rand" endpoint result from a
// HTTP "OK" response.
func NewRandResultViewOK(body *RandResponseBody) *demoviews.RandResultView {
	v := &demoviews.RandResultView{
		Result: body.Result,
	}

	return v
}

// NewRandInvalidArguments builds a demo service rand endpoint
// invalid_arguments error.
func NewRandInvalidArguments(body *RandInvalidArgumentsResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// ValidateRandInvalidArgumentsResponseBody runs the validations defined on
// rand_invalid_arguments_response_body
func ValidateRandInvalidArgumentsResponseBody(body *RandInvalidArgumentsResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}
