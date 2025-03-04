package logger

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerField zap.Field

func convertLoggerFieldsToZapFields(fields ...LoggerField) []zap.Field {
	// Initialize fields.
	var zapFields []zap.Field = make([]zapcore.Field, len(fields))

	// Convert each.
	for i := range fields {
		zapFields[i] = zap.Field(fields[i])
	}

	return zapFields
}

func String(key string, value string) LoggerField {
	return LoggerField(zap.String(key, value))
}

func Stringp(key string, value *string) LoggerField {
	return LoggerField(zap.Stringp(key, value))
}

func Strings(key string, value []string) LoggerField {
	return LoggerField(zap.Strings(key, value))
}

func ByteString(key string, value []byte) LoggerField {
	return LoggerField(zap.ByteString(key, value))
}

func Any(key string, value any) LoggerField {
	return LoggerField(zap.Any(key, value))
}

func Int(key string, value int) LoggerField {
	return LoggerField(zap.Int(key, value))
}

func Intp(key string, value *int) LoggerField {
	return LoggerField(zap.Intp(key, value))
}

func Int64(key string, value int64) LoggerField {
	return LoggerField(zap.Int64(key, value))
}

func Float64(key string, value float64) LoggerField {
	return LoggerField(zap.Float64(key, value))
}

func Error(value error) LoggerField {
	return LoggerField(zap.Error(value))
}

func Time(key string, value time.Time) LoggerField {
	return LoggerField(zap.Time(key, value))
}

func Timep(key string, value *time.Time) LoggerField {
	return LoggerField(zap.Timep(key, value))
}

func Bool(key string, value bool) LoggerField {
	return LoggerField(zap.Bool(key, value))
}

func Array(key string, value zapcore.ArrayMarshaler) LoggerField {
	return LoggerField(zap.Array(key, value))
}

func Errors(key string, value []error) LoggerField {
	return LoggerField(zap.Errors(key, value))
}

func Boolp(key string, value *bool) LoggerField {
	return LoggerField(zap.Boolp(key, value))
}
