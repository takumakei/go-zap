package loggers

import "go.uber.org/zap"

var nop = zap.NewNop()

// Nop returns always the same instance of no-op Logger.
func Nop() *zap.Logger {
	return nop
}

// L returns the first non-nil Logger in loggers or no-op Logger, never returns nil.
func L(loggers ...*zap.Logger) *zap.Logger {
	for _, logger := range loggers {
		if logger != nil {
			return logger
		}
	}
	return nop
}
