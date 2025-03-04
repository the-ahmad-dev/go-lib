package logger

import (
	"log"
	"time"

	"github.com/the-ahmad-dev/go-lib/random"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerC struct {
	*zap.Logger
}

var defaultLogger *LoggerC

func init() {
	defaultLogger = Init("logger")
}

func Init(app string) *LoggerC {
	// Configure.
	logConf := zap.NewProductionConfig()
	logConf.EncoderConfig.LevelKey = "l"
	logConf.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)

	// Build.
	l, err := logConf.Build()
	if err != nil {
		log.Panicf("failed to build logger: %v", err)
	}

	// Prepare.
	logger := l.With(
		zap.String("app", app),
		zap.String("apphash", random.Base62Str(9)),
	)
	logger = logger.WithOptions(
		zap.AddCallerSkip(1),
	)

	// Ensure all standard logs go through this logger now.
	zap.RedirectStdLog(logger)

	return &LoggerC{
		Logger: logger,
	}
}
