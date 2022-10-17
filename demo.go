package demoapi

import (
	"context"
	"log"

	demo "github.com/joberly/demo-go-api/gen/demo"
)

// demo service example implementation.
// The example methods log the requests and return zero values.
type demosrvc struct {
	logger *log.Logger
}

// NewDemo returns the demo service implementation.
func NewDemo(logger *log.Logger) demo.Service {
	return &demosrvc{logger}
}

// Rand implements rand.
func (s *demosrvc) Rand(ctx context.Context, p *demo.RandPayload) (res *demo.RandResult, err error) {
	res = &demo.RandResult{}
	s.logger.Print("demo.rand")
	return
}
