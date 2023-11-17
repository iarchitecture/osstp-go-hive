package core_logger

import (
	"go.uber.org/zap"
)

var logger *zap.Logger

func NewLoggerServer(name ...string) {
	logger = NewLogger(
		SetAppName("go-kit"),
		SetDevelopment(true),
		SetLevel(zap.DebugLevel),
	)
}

func GetLogger() *zap.Logger {
	return logger
}
