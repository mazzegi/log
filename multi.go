package log

import "github.com/mazzegi/log/entry"

type MultiWriter struct {
	targets []Writer
}

func NewMultiWriter(targets ...Writer) *MultiWriter {
	return &MultiWriter{
		targets: targets,
	}
}

func (mw *MultiWriter) Write(e entry.Entry) {
	for _, t := range mw.targets {
		t.Write(e)
	}
}
