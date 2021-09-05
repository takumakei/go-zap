go-zap
======================================================================

Utility packages for go.uber.org/zap.


loggers
----------------------------------------------------------------------

[![Go Reference](https://pkg.go.dev/badge/github.com/takumakei/go-zap/loggers.svg)](https://pkg.go.dev/github.com/takumakei/go-zap/loggers)

```
package loggers // import "github.com/takumakei/go-zap/loggers"

Package loggers centrally manages *zap.Logger throughout and across
packages.

FUNCTIONS

func L(loggers ...*zap.Logger) *zap.Logger
    L returns the first non-nil Logger in loggers or no-op Logger, never returns
    nil.

func Names() []string
    Names returns names registered.

func Nop() *zap.Logger
    Nop returns always the same instance of no-op Logger.

func Register(setter Setter)
    Register registers the setter function by its package name. It panics if the
    package name is already registered.

func RegisterName(setter Setter, name string)
    RegisterName registers the setter function by name. It panics if name is
    already registered.

func ResetLevel(enab zapcore.LevelEnabler) zap.Option
    ResetLevel returns an option that resets the LevelEnabler of the core.

func SetLogger(logger *zap.Logger, names ...string) error
    SetLogger calls registered setter functions by name. Each name may have a
    log level suffix like '=debug'.


TYPES

type Setter func(*zap.Logger)
    Setter is type of the function that the package receives the logger.

```
