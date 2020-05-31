package log

import (
	"fmt"
	"time"

	"github.com/mazzegi/log/entry"
)

type Writer interface {
	Write(e entry.Entry)
	Close()
}

type NamedLogger struct {
	name   string
	writer Writer
}

func NewNamed(name string, w Writer) *NamedLogger {
	l := &NamedLogger{
		name:   name,
		writer: w,
	}
	return l
}

func (l *NamedLogger) Close() {
	l.writer.Close()
}

func (l *NamedLogger) Logf(level entry.Level, s string, args ...interface{}) {
	l.writer.Write(entry.Entry{
		Time:      time.Now(),
		Level:     level,
		Component: l.name,
		Message:   fmt.Sprintf(s, args...),
	})
}
