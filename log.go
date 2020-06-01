package log

import (
	"fmt"

	"github.com/mazzegi/log/console"
	"github.com/mazzegi/log/entry"
)

func Debugf(s string, args ...interface{}) {
	logger.Log(entry.Make(entry.LevelDebug, s, args...))
}

func Infof(s string, args ...interface{}) {
	logger.Log(entry.Make(entry.LevelInfo, s, args...))
}

func Warnf(s string, args ...interface{}) {
	logger.Log(entry.Make(entry.LevelWarn, s, args...))
}

func Errorf(s string, args ...interface{}) {
	logger.Log(entry.Make(entry.LevelError, s, args...))
}

func Fatalf(s string, args ...interface{}) {
	logger.Log(entry.Make(entry.LevelFatal, s, args...))
	panic(fmt.Sprintf(s, args...))
}

//

type Logger interface {
	Log(e entry.Entry)
}

var logger Logger = NewStdLogger("default", console.NewWriter())

func Install(l Logger) {
	logger = l
}
