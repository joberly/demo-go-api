package log

import (
	"fmt"

	"go.uber.org/zap"
	"goa.design/goa/v3/middleware"
)

type zapAdapter struct {
	logger *zap.Logger
}

func NewAdapter(logger *zap.Logger) middleware.Logger {
	// Make a buffer of zap.Field that should be more than enough.
	// Log will expand it as necessary.
	return &zapAdapter{
		logger: logger,
	}
}

func (z *zapAdapter) Log(keyvals ...interface{}) error {
	// TODO use a sync.Pool to use an preallocated array
	fields := make([]zap.Field, len(keyvals))
	for i := 0; i < len(keyvals); i += 2 {
		fields[i/2] = zap.Any(fmt.Sprint(keyvals[i]), keyvals[i+1])
	}
	z.logger.Info("log", fields...)
	return nil
}
