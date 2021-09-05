package loggers

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type levels struct {
	logger *zap.Logger
	cache  map[zapcore.Level]*zap.Logger
}

func newLevels(logger *zap.Logger) *levels {
	return &levels{
		logger: logger,
		cache:  make(map[zapcore.Level]*zap.Logger),
	}
}

func (l levels) Get(lvl zapcore.Level) *zap.Logger {
	if eq(l.logger.Core(), lvl) {
		return l.logger
	}
	if v, ok := l.cache[lvl]; ok {
		return v
	}
	v := l.logger.WithOptions(ResetLevel(lvl))
	l.cache[lvl] = v
	return v
}

func eq(lhs, rhs zapcore.LevelEnabler) bool {
	for i := int8(zapcore.DebugLevel); i <= int8(zapcore.FatalLevel); i++ {
		lvl := zapcore.Level(i)
		if lhs.Enabled(lvl) != rhs.Enabled(lvl) {
			return false
		}
	}
	return true
}
