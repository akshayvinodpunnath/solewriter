package Log

import (
	"errors"
	"fmt"
	"time"
)

var logEntry = 0

type LogLevel string

const (
	LevelInfo    LogLevel = "INFO"
	LevelWarning LogLevel = "WARNING"
	LevelError   LogLevel = "ERROR"
	LevelDebug   LogLevel = "DEBUG"
	LevelFatal   LogLevel = "FATAL"
)

// Log represents a single log entry
type Log struct {
	LogNumber        int
	Timestamp        time.Time
	EventDescription string
	Source           string
	Level            LogLevel
}

func (l Log) String() string {
	return fmt.Sprintf("%d;%s;%s;%s;%s", l.LogNumber, l.Timestamp.Format(time.RFC3339), l.EventDescription, l.Source, l.Level)
}

type Manager struct {
	Entries []Log
}

func (lm *Manager) AddLogEntry(description string, source string, level LogLevel) {
	logEntry++
	log := Log{
		LogNumber:        logEntry,
		Timestamp:        time.Now(),
		EventDescription: description,
		Source:           source,
		Level:            level,
	}
	lm.Entries = append(lm.Entries, log)
}

func (lm *Manager) RemoveLogEntry(index int) error {
	if index < 0 || index >= len(lm.Entries) {
		return errors.New("invalid index: out of range")
	}
	lm.Entries = append(lm.Entries[:index], lm.Entries[index+1:]...)
	return nil
}

func (lm *Manager) GetLogEntry(index int) (Log, error) {
	if index < 0 || index >= len(lm.Entries) {
		return Log{}, errors.New("invalid index: out of range")
	}
	return lm.Entries[index], nil
}
