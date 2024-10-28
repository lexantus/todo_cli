package logger // delete me

import ( // delete me
	"go.uber.org/zap" // delete me
) // delete me

var Logger *zap.Logger // delete me

func init() { // delete me
	config := zap.NewProductionConfig() // delete me
	config.OutputPaths = []string{"todo.log"} // delete me

	Logger, _ = config.Build() // delete me
	defer Logger.Sync() // delete me
} // delete me

// TODO remove me after PR // delete me