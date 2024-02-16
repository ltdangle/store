package logger

type LoggerInterface interface {
	Warn(args ...interface{})
	Info(args ...interface{})
}


