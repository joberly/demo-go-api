package demoapi

import (
	"context"
	"fmt"
	"math"
	"math/rand"
	"time"

	demo "github.com/joberly/demo-go-api/gen/demo"
	"go.uber.org/zap"
)

// demo service example implementation.
// The example methods log the requests and return zero values.
type demosrvc struct {
	rand   *rand.Rand
	logger *zap.Logger
}

// NewDemo returns the demo service implementation.
func NewDemo(logger *zap.Logger) demo.Service {
	return &demosrvc{
		rand:   rand.New(rand.NewSource(time.Now().UnixMicro())),
		logger: logger,
	}
}

// Rand implements rand.
func (s *demosrvc) Rand(ctx context.Context, p *demo.RandPayload) (*demo.RandResult, error) {
	s.logger.Info("demo.rand")

	var min int64 = 0
	var max int64 = math.MaxInt64
	var numErrors = 0
	if p.Max != nil {
		max = *p.Max
		if max <= 0 {
			numErrors++
			s.logger.Error("max must be greater than zero", zap.Int64("max", max))
		}
	}
	if p.Min != nil {
		min = *p.Min
		if min < 0 {
			numErrors++
			s.logger.Error("min must be greater than or equal to zero", zap.Int64("min", min))
		}
	}
	if max <= min {
		numErrors++
		s.logger.Error("max is not greater than min",
			zap.Int64("min", min),
			zap.Int64("max", max),
		)
	}

	if numErrors > 0 {
		return nil, fmt.Errorf("invalid argument(s)")
	}

	res := demo.RandResult{
		Result: new(int64),
	}
	*res.Result = s.rand.Int63n(max-min) + min

	return &res, nil
}
