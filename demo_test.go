package demoapi

import (
	"context"
	"math"
	"testing"

	demo "github.com/joberly/demo-go-api/gen/demo"
	"github.com/openlyinc/pointy"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

type randTest struct {
	payload    demo.RandPayload
	expSuccess bool
}

func TestRand(t *testing.T) {
	ctx := context.TODO()
	matrix := []randTest{
		{
			payload: demo.RandPayload{
				Min: pointy.Int64(0),
				Max: pointy.Int64(math.MaxInt64),
			},
			expSuccess: true,
		},
		{
			payload: demo.RandPayload{
				Min: pointy.Int64(-1),
				Max: pointy.Int64(math.MaxInt64 / 2),
			},
			expSuccess: false,
		},
		{
			payload: demo.RandPayload{
				Min: pointy.Int64(math.MaxInt64 / 2),
				Max: pointy.Int64(math.MinInt64 / 2),
			},
			expSuccess: false,
		},
		{
			payload: demo.RandPayload{
				Min: nil,
				Max: nil,
			},
			expSuccess: true,
		},
		{
			payload: demo.RandPayload{
				Min: nil,
				Max: pointy.Int64(math.MaxInt64 / 2),
			},
			expSuccess: true,
		},
		{
			payload: demo.RandPayload{
				Min: pointy.Int64(math.MaxInt64 / 2),
				Max: nil,
			},
			expSuccess: true,
		},
		{
			payload: demo.RandPayload{
				Min: pointy.Int64(1),
				Max: pointy.Int64(10),
			},
			expSuccess: true,
		},
	}

	logger := zap.NewNop()
	srvc := NewDemo(logger)
	ta := assert.New(t)
	for _, r := range matrix {
		res, err := srvc.Rand(ctx, &r.payload)
		if r.expSuccess {
			min := pointy.Int64Value(r.payload.Min, 0)
			max := pointy.Int64Value(r.payload.Max, math.MaxInt64)
			ta.NotNil(res)
			ta.NoError(err)
			if res != nil {
				ta.NotNil(res.Result)
				if res.Result != nil {
					resVal := *res.Result
					ta.GreaterOrEqual(resVal, min)
					ta.Less(resVal, max)
				}
			}
		} else {
			ta.Nil(res)
			ta.Error(err)
		}
	}
}
