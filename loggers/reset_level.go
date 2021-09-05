package loggers

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ResetLevel returns an option that resets the LevelEnabler of the core.
func ResetLevel(enab zapcore.LevelEnabler) zap.Option {
	return zap.WrapCore(func(core zapcore.Core) zapcore.Core {
		return newLevelFreeCore(core, enab)
	})
}

func newLevelFreeCore(core zapcore.Core, enab zapcore.LevelEnabler) zapcore.Core {
	if v, ok := core.(*levelFreeCore); ok {
		return &levelFreeCore{Core: v.Core, enab: enab}
	}
	return &levelFreeCore{Core: core, enab: enab}
}

type levelFreeCore struct {
	zapcore.Core
	enab zapcore.LevelEnabler
}

func (c *levelFreeCore) Enabled(lvl zapcore.Level) bool {
	return c.enab.Enabled(lvl)
}

func (c *levelFreeCore) Check(ent zapcore.Entry, ce *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if c.Enabled(ent.Level) {
		return ce.AddCore(ent, c)
	}
	return ce
}
