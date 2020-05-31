package log

import (
	"fmt"

	"github.com/mazzegi/log/console"
	"github.com/mazzegi/log/entry"
)

func Debugf(s string, args ...interface{}) {
	logger.Logf(entry.LevelDebug, s, args...)
}

func Infof(s string, args ...interface{}) {
	logger.Logf(entry.LevelInfo, s, args...)
}

func Warnf(s string, args ...interface{}) {
	logger.Logf(entry.LevelWarn, s, args...)
}

func Errorf(s string, args ...interface{}) {
	logger.Logf(entry.LevelError, s, args...)
}

func Fatalf(s string, args ...interface{}) {
	logger.Logf(entry.LevelFatal, s, args...)
	panic(fmt.Sprintf(s, args...))
}

//

type Logger interface {
	Logf(level entry.Level, s string, args ...interface{})
}

var logger Logger = NewNamed("default", console.NewWriter())

func Install(l Logger) {
	logger = l
}
