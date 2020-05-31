package console

import (
	"io"
	"os"

	"github.com/mazzegi/log/entry"
)

type Formatter interface {
	Format(entry.Entry) string
}

type Option func(w *Writer)

func WithStream(s io.Writer) Option {
	return func(w *Writer) {
		w.writer = s
	}
}

func WithFormatter(f Formatter) Option {
	return func(w *Writer) {
		w.formatter = f
	}
}

type Writer struct {
	writer    io.Writer
	formatter Formatter
}

func NewWriter(opts ...Option) *Writer {
	w := &Writer{
		writer:    os.Stdout,
		formatter: ColorFormatter{},
	}
	for _, o := range opts {
		o(w)
	}
	return w
}

func (w *Writer) Write(e entry.Entry) {
	w.writer.Write([]byte(w.formatter.Format(e) + "\n"))
}
