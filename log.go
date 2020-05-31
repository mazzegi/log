package log

import (
	"fmt"
	"time"

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
type Writer interface {
	Write(e entry.Entry)
}

var logger = NewLogger("default")

func InstallGlobal(l *Logger) {
	logger = l
}

type Option func(l *Logger)

func WithWriter(w Writer) Option {
	return func(l *Logger) {
		l.writer = w
	}
}

type Logger struct {
	writer Writer
	name   string
}

func NewLogger(name string, opts ...Option) *Logger {
	l := &Logger{
		writer: console.NewWriter(),
		name:   name,
	}
	for _, opt := range opts {
		opt(l)
	}
	return l
}

func (l *Logger) Logf(level entry.Level, s string, args ...interface{}) {
	l.writer.Write(entry.Entry{
		Time:      time.Now(),
		Level:     level,
		Component: l.name,
		Message:   fmt.Sprintf(s, args...),
	})
}
