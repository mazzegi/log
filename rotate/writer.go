package rotate

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/mazzegi/log/entry"
	"github.com/pkg/errors"
)

const (
	KB int = 1024
	MB int = KB * 1024
	GB int = MB * 1024
)

type Option func(w *Writer) error

func WithFileSize(s int) Option {
	return func(w *Writer) error {
		if s < KB || s > 10*GB {
			return errors.Errorf("invalid file-size: %d", s)
		}
		w.fileSize = s
		return nil
	}
}

func WithFileCount(c int) Option {
	return func(w *Writer) error {
		if c < 2 {
			return errors.Errorf("invalid file-count: %d", c)
		}
		w.fileCount = c
		return nil
	}
}

func WithFileOpErrHandler(h func(err error)) Option {
	return func(w *Writer) error {
		w.onFileOpErr = h
		return nil
	}
}

type Writer struct {
	file        io.WriteCloser
	fileDir     string
	fileBase    string
	fileSize    int
	fileCount   int
	currSize    int
	onFileOpErr func(error)
	msgC        chan string
	doneC       chan struct{}
}

func NewWriter(path string, opts ...Option) (*Writer, error) {
	err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
	if err != nil {
		return nil, err
	}
	f, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	w := &Writer{
		file:        f,
		fileDir:     filepath.Dir(path),
		fileBase:    filepath.Base(path),
		fileSize:    1 * MB,
		fileCount:   5,
		currSize:    0,
		onFileOpErr: func(err error) { panic(err) },
		msgC:        make(chan string, 10000),
		doneC:       make(chan struct{}),
	}
	for _, opt := range opts {
		err := opt(w)
		if err != nil {
			return nil, err
		}
	}
	go func() {
		defer close(w.doneC)
		for s := range w.msgC {
			n, _ := w.file.Write([]byte(s + "\n"))
			w.currSize += n
			if w.currSize >= w.fileSize {
				w.rotate()
			}
		}
	}()
	return w, nil
}

func (w *Writer) Close() {
	close(w.msgC)
	<-w.doneC
	w.file.Close()
}

func (w *Writer) Write(e entry.Entry) {
	w.msgC <- fmt.Sprintf("%s [%s] [%s] [%s] %s", e.Time.Format("2006-01-02T15:04:05.000"), e.Program, e.Component, e.Level, e.Message)
}

func (w *Writer) rotate() {
	handleErr := func(err error) {
		if err != nil {
			w.onFileOpErr(err)
		}
	}
	w.file.Close()
	pattern := filepath.Join(w.fileDir, w.fileBase+".*")
	matches, _ := filepath.Glob(pattern)

	//remove oldest
	for len(matches) > w.fileCount {
		last := matches[len(matches)-1]
		err := os.Remove(last)
		handleErr(err)
		matches = matches[:len(matches)-1]
	}

	//rename others
	for i := len(matches) - 1; i >= 0; i-- {
		new := filepath.Join(w.fileDir, w.fileBase+fmt.Sprintf(".%03d", i+2))
		err := os.Rename(matches[i], new)
		handleErr(err)
	}
	err := os.Rename(filepath.Join(w.fileDir, w.fileBase), filepath.Join(w.fileDir, w.fileBase+".001"))
	handleErr(err)

	f, err := os.OpenFile(filepath.Join(w.fileDir, w.fileBase), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	handleErr(err)
	w.file = f
	w.currSize = 0
}
