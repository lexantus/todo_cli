package logger

import (
	"go.uber.org/zap"
)

var Logger *zap.Logger

func init() {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"todo.log"}

	Logger, _ = config.Build()
	defer Logger.Sync()
}

func Error(msg string, fields ...zap.Field) {
	Logger.Error(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	Logger.Info(msg, fields...)
}
