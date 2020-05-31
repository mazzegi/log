package entry

import (
	"fmt"
	"time"
)

type Level string

const (
	LevelDebug Level = "DEBUG"
	LevelInfo  Level = "INFO "
	LevelWarn  Level = "WARN "
	LevelError Level = "ERROR"
	LevelFatal Level = "FATAL"
)

type Entry struct {
	Time      time.Time
	Level     Level
	Program   string
	Component string
	Message   string
}

func Make(level Level, s string, args ...interface{}) Entry {
	return Entry{
		Time:    time.Now(),
		Level:   level,
		Message: fmt.Sprintf(s, args...),
	}
}
