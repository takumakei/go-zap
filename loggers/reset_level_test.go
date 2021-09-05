package loggers_test

import (
	"github.com/takumakei/go-zap/loggers"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func ExampleResetLevel() {
	logWarn := zap.NewExample(
		zap.IncreaseLevel(zap.WarnLevel),
	)

	logInfo := logWarn.WithOptions(
		loggers.ResetLevel(zapcore.DebugLevel),
	)

	logWarn.Info("info")
	logWarn.Warn("warn")
	logInfo.Info("info")
	logInfo.Warn("warn")

	// Output:
	// {"level":"warn","msg":"warn"}
	// {"level":"info","msg":"info"}
	// {"level":"warn","msg":"warn"}
}
