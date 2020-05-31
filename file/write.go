package file

import (
	"fmt"
	"io"
	"os"

	"github.com/mazzegi/log/entry"
)

type Writer struct {
	file io.WriteCloser
}

func NewWriter(path string) (*Writer, error) {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	return &Writer{
		file: f,
	}, nil
}

func (w *Writer) Close() {
	w.file.Close()
}

func (w *Writer) Write(e entry.Entry) {
	fmt.Fprintf(w.file, "%s [%s] [%s] %s\n", e.Time.Format("2006-01-02T15:04:05.000"), e.Component, e.Level, e.Message)
}
