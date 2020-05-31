package log

import (
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

func (l *NamedLogger) Log(e entry.Entry) {
	if e.Program == "" {
		e.Program = l.name
	}
	l.writer.Write(e)
}
