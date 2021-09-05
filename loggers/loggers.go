// Package loggers centrally manages *zap.Logger throughout and across packages.
package loggers

import (
	"fmt"
	"strings"

	"github.com/takumakei/go-funcname"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Setter is type of the function that the package receives the logger.
type Setter func(*zap.Logger)

var (
	setters = make(map[string]Setter)
	names   []string
)

// Register registers the setter function by its package name.
// It panics if the package name is already registered.
func Register(setter Setter) {
	name, pkgname := funcname.Of(setter)
	if _, ok := setters[pkgname]; ok {
		panic(fmt.Sprintf("'%s(.%s)' has already been registered", pkgname, name))
	}
	setters[pkgname] = setter
	names = append(names, pkgname)
}

// RegisterName registers the setter function by name.
// It panics if name is already registered.
func RegisterName(setter Setter, name string) {
	if _, ok := setters[name]; ok {
		panic(fmt.Sprintf("'%s' has already been registered", name))
	}
	setters[name] = setter
	names = append(names, name)
}

// Names returns names registered.
func Names() []string {
	return append(make([]string, 0, len(names)), names...)
}

// SetLogger calls registered setter functions by name.
// Each name may have a log level suffix like '=debug'.
func SetLogger(logger *zap.Logger, names ...string) error {
	levels := newLevels(logger)
	for _, s := range names {
		name, level := splitName(s)
		setter, ok := setters[name]
		if !ok {
			return fmt.Errorf("not registered %q", name)
		}
		if len(level) > 0 {
			var lvl zapcore.Level
			if err := lvl.Set(level); err != nil {
				return fmt.Errorf("%w %q(%s)", err, level, name)
			}
			setter(levels.Get(lvl))
		} else {
			setter(logger)
		}
	}
	return nil
}

func splitName(s string) (name, level string) {
	if i := strings.IndexByte(s, '='); i == -1 {
		name = s
	} else {
		name, level = s[:i], s[i+1:]
	}
	return
}
