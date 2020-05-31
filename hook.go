package log

import (
	"fmt"

	"github.com/mazzegi/log/entry"
)

type Hook struct {
	hookFnc func(e entry.Entry) entry.Entry
}

func NewHook(hookFnc func(e entry.Entry) entry.Entry) *Hook {
	return &Hook{
		hookFnc: hookFnc,
	}
}

func (lh *Hook) Debugf(s string, args ...interface{}) {
	e := entry.Make(entry.LevelDebug, s, args...)
	logger.Log(lh.hookFnc(e))
}

func (lh *Hook) Infof(s string, args ...interface{}) {
	e := entry.Make(entry.LevelInfo, s, args...)
	logger.Log(lh.hookFnc(e))
}

func (lh *Hook) Warnf(s string, args ...interface{}) {
	e := entry.Make(entry.LevelWarn, s, args...)
	logger.Log(lh.hookFnc(e))
}

func (lh *Hook) Errorf(s string, args ...interface{}) {
	e := entry.Make(entry.LevelError, s, args...)
	logger.Log(lh.hookFnc(e))
}

func (lh *Hook) Fatalf(s string, args ...interface{}) {
	e := entry.Make(entry.LevelFatal, s, args...)
	logger.Log(lh.hookFnc(e))
	panic(fmt.Sprintf(s, args...))
}

func ComponentHook(comp string) *Hook {
	return NewHook(func(e entry.Entry) entry.Entry {
		e.Component = comp
		return e
	})
}
