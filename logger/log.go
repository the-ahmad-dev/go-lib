package logger

func (c *LoggerC) Debug(msg string, fields ...LoggerField) {
	c.Logger.Debug(msg, convertLoggerFieldsToZapFields(fields...)...)
}

func (c *LoggerC) Info(msg string, fields ...LoggerField) {
	c.Logger.Info(msg, convertLoggerFieldsToZapFields(fields...)...)
}

func (c *LoggerC) Warn(msg string, fields ...LoggerField) {
	c.Logger.Warn(msg, convertLoggerFieldsToZapFields(fields...)...)
}

func (c *LoggerC) Error(msg string, fields ...LoggerField) {
	c.Logger.Error(msg, convertLoggerFieldsToZapFields(fields...)...)
}

func (c *LoggerC) Fatal(msg string, fields ...LoggerField) {
	c.Logger.Fatal(msg, convertLoggerFieldsToZapFields(fields...)...)
}

func (c *LoggerC) Panic(msg string, fields ...LoggerField) {
	c.Logger.Panic(msg, convertLoggerFieldsToZapFields(fields...)...)
}
