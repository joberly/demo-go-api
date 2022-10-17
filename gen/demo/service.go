// Code generated by goa v3.10.1, DO NOT EDIT.
//
// demo service
//
// Command:
// $ goa gen github.com/joberly/demo-go-api/design

package demo

import (
	"context"

	demoviews "github.com/joberly/demo-go-api/gen/demo/views"
	goa "goa.design/goa/v3/pkg"
)

// Demonstration Go Service
type Service interface {
	// Rand implements rand.
	Rand(context.Context, *RandPayload) (res *RandResult, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "demo"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"rand"}

// RandPayload is the payload type of the demo service rand method.
type RandPayload struct {
	Min *int
	Max *int
}

// RandResult is the result type of the demo service rand method.
type RandResult struct {
	Result *int
}

// MakeInvalidArguments builds a goa.ServiceError from an error.
func MakeInvalidArguments(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "invalid_arguments", false, false, false)
}

// NewRandResult initializes result type RandResult from viewed result type
// RandResult.
func NewRandResult(vres *demoviews.RandResult) *RandResult {
	return newRandResult(vres.Projected)
}

// NewViewedRandResult initializes viewed result type RandResult from result
// type RandResult using the given view.
func NewViewedRandResult(res *RandResult, view string) *demoviews.RandResult {
	p := newRandResultView(res)
	return &demoviews.RandResult{Projected: p, View: "default"}
}

// newRandResult converts projected type RandResult to service type RandResult.
func newRandResult(vres *demoviews.RandResultView) *RandResult {
	res := &RandResult{
		Result: vres.Result,
	}
	return res
}

// newRandResultView projects result type RandResult to projected type
// RandResultView using the "default" view.
func newRandResultView(res *RandResult) *demoviews.RandResultView {
	vres := &demoviews.RandResultView{
		Result: res.Result,
	}
	return vres
}