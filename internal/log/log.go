package log

import (
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
)

var (
	logger *zap.SugaredLogger
)

func init() {
	l, _ := zap.NewProduction()
	grpc_zap.ReplaceGrpcLogger(l)
	logger = l.Sugar()
}

// Fatal logs a message and calls os.Exit
func Fatal(err error) {
	logger.Fatalw(err.Error())
}

// Info logs an informational message
func Info(message string, keysAndValues ...interface{}) {
	if len(keysAndValues) == 0 {
		logger.Info(message)
	} else {
		logger.Infow(message, keysAndValues...)
	}
}
