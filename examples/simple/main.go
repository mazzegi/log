package main

import (
	"fmt"
	"os"

	"github.com/mazzegi/log"
	"github.com/mazzegi/log/console"
	"github.com/mazzegi/log/entry"
	"github.com/mazzegi/log/file"
)

func createLogger() *log.NamedLogger {
	w := console.NewWriter(
		console.WithFormatter(console.ColorFormatter{}),
		console.WithStream(os.Stderr),
	)

	ew, err := file.NewWriter("error.log")
	if err != nil {
		panic(err)
	}
	ef := log.NewFilter(
		func(e entry.Entry) bool {
			return e.Level == entry.LevelError ||
				e.Level == entry.LevelFatal
		},
		ew,
	)

	return log.NewNamed(
		"simple-example",
		log.NewMultiWriter(w, ef),
	)
}

func main() {
	l := createLogger()
	log.Install(l)
	defer l.Close()

	log.Debugf("a simple debug message %d", 42)
	log.Infof("a simple info message %q", "dude")
	log.Warnf("a simple warn message %T", true)
	log.Errorf("a simple error message %f", 1.23)
	log.Fatalf("finally ... %v", fmt.Errorf("an expected error occurred"))
}
