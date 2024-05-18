package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type Logger struct {
	Logger    *log.Logger
	formatter Formatter
	debugMode bool
	prefix    string
}

var logger *Logger

func InitializeLogger(prefix string, debugMode bool) *Logger {
	flags := log.Ldate | log.Ltime
	logger = &Logger{
		Logger:    log.New(os.Stdout, prefix+" ", flags),
		formatter: &JSONFormatter{},
		debugMode: debugMode,
		prefix:    prefix,
	}

	return logger
}

func GetLogger() *Logger {
	return logger
}

func (l *Logger) logWithLevel(level LogLevel, message string, extra interface{}) {
	if !l.debugMode && level == DEBUG {
		return
	}

	timestamp := time.Now()

	l.logInLoki(message, timestamp, level, extra)
	l.logInConsole(message, timestamp, level, extra)
}

func (l *Logger) logInLoki(message string, timestamp time.Time, level LogLevel, extra interface{}) {

	LOKI_URL := "http://loki:3100/loki/api/v1/push"

	data := map[string]interface{}{
		"streams": []map[string]interface{}{
			{
				"stream": map[string]string{
					"job":      l.prefix,
					"level":    logLevelString(level),
					"filename": getFilePath(),
				},
				"values": [][]interface{}{
					{
						fmt.Sprintf("%d", timestamp.UnixNano()),
						l.generateLogMessage(message, timestamp, level, extra),
					},
				},
			},
		},
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error al codificar los datos JSON:", err)
		return
	}

	req, err := http.NewRequest("POST", LOKI_URL, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error al crear la solicitud HTTP:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error al hacer la solicitud HTTP:", err)
		return
	}

	defer resp.Body.Close()
}

func (l *Logger) logInConsole(message string, timestamp time.Time, level LogLevel, extra interface{}) {
	l.Logger.Println(levelColorFunc[level](l.generateLogMessage(message, timestamp, level, extra)))
}

func (l *Logger) generateLogMessage(message string, _ time.Time, level LogLevel, extra interface{}) string {
	return fmt.Sprintf("[%s] %s - extra: %v", logLevelString(level), message, extra)
}
