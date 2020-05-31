package main

import (
	"os"

	"github.com/mazzegi/log"
	"github.com/mazzegi/log/console"
	"github.com/mazzegi/log/entry"
)

func setupLogger() {
	w := console.NewWriter(
		console.WithFormatter(console.ColorFormatter{}),
		console.WithStream(os.Stderr),
	)

	ef := log.NewFilter(
		func(e entry.Entry) bool {
			return e.Level == entry.LevelError
		},
		console.NewWriter(
			console.WithFormatter(console.ColorFormatter{}),
			console.WithStream(os.Stdout),
		),
	)

	log.InstallGlobal(log.NewLogger(
		"simple-example",
		log.WithWriter(
			log.NewMultiWriter(w, ef),
		),
	))
}

func main() {
	setupLogger()

	log.Debugf("a simple debug message %d", 42)
	log.Infof("a simple info message %q", "dude")
	log.Warnf("a simple warn message %T", true)
	log.Errorf("a simple error message %f", 1.23)
	//log.Fatalf("finally ... %v", fmt.Errorf("an expected error occurred"))
}
