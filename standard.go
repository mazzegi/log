package log

import (
	"github.com/mazzegi/log/entry"
)

type Writer interface {
	Write(e entry.Entry)
	Close()
}

type StdLogger struct {
	name   string
	writer Writer
}

func NewStdLogger(name string, w Writer) *StdLogger {
	l := &StdLogger{
		name:   name,
		writer: w,
	}
	return l
}

func (l *StdLogger) Close() {
	l.writer.Close()
}

func (l *StdLogger) Log(e entry.Entry) {
	if e.Program == "" {
		e.Program = l.name
	}
	l.writer.Write(e)
}
