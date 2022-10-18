package demoapi

import (
	"context"

	demo "github.com/joberly/demo-go-api/gen/demo"
	"go.uber.org/zap"
)

// demo service example implementation.
// The example methods log the requests and return zero values.
type demosrvc struct {
	logger *zap.Logger
}

// NewDemo returns the demo service implementation.
func NewDemo(logger *zap.Logger) demo.Service {
	return &demosrvc{logger}
}

// Rand implements rand.
func (s *demosrvc) Rand(ctx context.Context, p *demo.RandPayload) (res *demo.RandResult, err error) {
	res = &demo.RandResult{}
	s.logger.Info("demo.rand")
	return
}
