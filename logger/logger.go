package logger

func Info(args ...interface{}) {
	logger.logger.Info(args...)
}

func Warn(args ...interface{}) {
	logger.logger.Warn(args...)
}

func Debug(args ...interface{}) {
	logger.logger.Debug(args...)
}

func Error(args ...interface{}) {
	logger.logger.Error(args...)
}

func Fatal(args ...interface{}) {
	logger.logger.Fatal(args...)
}

func Panic(args ...interface{}) {
	logger.logger.Panic(args...)
}

func Infof(format string, args ...interface{}) {
	logger.logger.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	logger.logger.Warnf(format, args...)
}

func Debugf(format string, args ...interface{}) {
	logger.logger.Debugf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	logger.logger.Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	logger.logger.Fatalf(format, args...)
}

func Panicf(format string, args ...interface{}) {
	logger.logger.Panicf(format, args...)
}
