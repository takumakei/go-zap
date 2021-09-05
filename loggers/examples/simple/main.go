package main

import (
	"github.com/takumakei/go-zap/loggers"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger *zap.Logger = zap.NewNop()
	logA   *zap.Logger = logger
	logB   *zap.Logger = logger
	logC   *zap.Logger = logger
	logD   *zap.Logger = logger
)

func setA(logger *zap.Logger) { logA = logger }
func setB(logger *zap.Logger) { logB = logger }
func setC(logger *zap.Logger) { logC = logger }
func setD(logger *zap.Logger) { logD = logger }

func main() {
	loggers.RegisterName(setA, "a")
	loggers.RegisterName(setB, "b")
	loggers.RegisterName(setC, "c")
	loggers.RegisterName(setD, "d")

	logger = zap.NewExample(zap.IncreaseLevel(zapcore.InfoLevel))

	if err := loggers.SetLogger(logger, "a=debug", "b=warn", "c"); err != nil {
		logger.Fatal("loggers.SetLogger", zap.Error(err))
	}

	// logger's level enabler is Info, so...
	logger.Debug("logger", zap.Bool("skip", true)) // skip
	logger.Info("logger", zap.Bool("skip", false)) // ok
	logger.Warn("logger", zap.Bool("skip", false)) // ok

	// logA's level enabler is debug, so...
	logA.Debug("logA", zap.Bool("skip", false)) // ok
	logA.Info("logA", zap.Bool("skip", false))  // ok
	logA.Warn("logA", zap.Bool("skip", false))  // ok

	// logB's level enabler is warn, so...
	logB.Debug("logB", zap.Bool("skip", true)) // skip
	logB.Info("logB", zap.Bool("skip", true))  // skip
	logB.Warn("logB", zap.Bool("skip", false)) // ok

	// logC's level enabler is same as logger's Info, so...
	logC.Debug("logC", zap.Bool("skip", true)) // skip
	logC.Info("logC", zap.Bool("skip", false)) // ok
	logC.Warn("logC", zap.Bool("skip", false)) // ok

	// logD's logger is Nop, so...
	logD.Debug("logD", zap.Bool("skip", true)) // skip
	logD.Info("logD", zap.Bool("skip", true))  // skip
	logD.Warn("logD", zap.Bool("skip", true))  // skip
}
