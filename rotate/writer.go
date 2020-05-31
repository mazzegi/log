package rotate

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/mazzegi/log/entry"
)

const (
	KB int = 1024
	MB int = KB * 1024
	GB int = MB * 1024
)

type Writer struct {
	file      io.WriteCloser
	fileDir   string
	fileBase  string
	fileSize  int
	fileCount int
	currSize  int
	msgC      chan string
	doneC     chan struct{}
}

func NewWriter(path string) (*Writer, error) {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	w := &Writer{
		file:      f,
		fileDir:   filepath.Dir(path),
		fileBase:  filepath.Base(path),
		fileSize:  1 * MB,
		fileCount: 5,
		currSize:  0,
		msgC:      make(chan string, 10000),
		doneC:     make(chan struct{}),
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
	w.file.Close()
	//tmp := filepath.Join(w.fileDir, w.fileBase+"."+uuid.MakeV4())
	//ioutil.ReadDir()
}
