package entry

import "time"

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
	Component string
	Message   string
}
