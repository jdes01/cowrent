package logger

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/fatih/color"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARNING
	ERROR
)

type Formatter interface {
	Format(interface{}) string
}

type JSONFormatter struct{}

func (jf *JSONFormatter) Format(extra interface{}) string {
	if extra == nil {
		return ""
	}
	data, err := json.Marshal(extra)
	if err != nil {
		return fmt.Sprintf(`{"message": "Error formatting extra: %v"}`, err)
	}
	return string(data)
}

func (l *Logger) Debug(message string, extra interface{}) {
	l.logWithLevel(DEBUG, message, extra)
}

func (l *Logger) Info(message string, extra interface{}) {
	l.logWithLevel(INFO, message, extra)
}

func (l *Logger) Warning(message string, extra interface{}) {
	l.logWithLevel(WARNING, message, extra)
}

func (l *Logger) Error(message string, extra interface{}) {
	l.logWithLevel(ERROR, message, extra)
}

func logLevelString(level LogLevel) string {
	switch level {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}

var levelColorFunc = map[LogLevel]func(...interface{}) string{
	DEBUG:   color.New(color.FgBlue).SprintFunc(),
	INFO:    color.New(color.FgGreen).SprintFunc(),
	WARNING: color.New(color.FgYellow).SprintFunc(),
	ERROR:   color.New(color.FgRed).SprintFunc(),
}

func getFilePath() string {
	pc, _, _, _ := runtime.Caller(4)

	filePath, err := filepath.Abs(runtime.FuncForPC(pc).Name())
	if err != nil {
		filePath = "unknown"
	}
	return filePath
}
