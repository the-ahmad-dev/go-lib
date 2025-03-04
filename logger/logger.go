package logger

import "context"

// Context key to pass logger.
var CtxKey = &struct{}{}

// Get returns the default logger.
func Get(ctx context.Context) *LoggerC {
	v := ctx.Value(CtxKey)
	if v == nil {
		return defaultLogger
	}

	logger, ok := v.(*LoggerC)
	if !ok {
		return defaultLogger
	}

	return logger
}

func (c *LoggerC) GetWithFields(fields ...LoggerField) *LoggerC {
	// Return same if no fields.
	if len(fields) == 0 {
		return c
	}

	// Create a new logger.
	new := *c

	// Add default fields to the logger.
	new.Logger = new.Logger.With(convertLoggerFieldsToZapFields(fields...)...)

	return &new
}
