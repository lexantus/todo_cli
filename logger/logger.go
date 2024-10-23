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

// TODO remove me after PR
