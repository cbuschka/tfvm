package log

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Level int

// Log levels.
const (
	TRACE Level = iota
	DEBUG
	INFO
	NONE
)

var levelNames = []string{
	"TRACE",
	"DEBUG",
	"INFO",
	"NONE",
}

var threshold = initLogging()

func getThresholdFromEnv() Level {
	logLevel := strings.TrimSpace(strings.ToUpper(os.Getenv("TFVM_LOG")))
	for index, name := range levelNames {
		if logLevel == name {
			return Level(index)
		}
	}

	return NONE
}

// Trace logs message to stderr.
func Trace(s string) {
	if IsTraceEnabled() {
		writeLog(time.Now(), TRACE, s)
	}
}

// Tracef logs message to stderr.
func Tracef(format string, args ...interface{}) {
	if IsTraceEnabled() {
		writeLogf(time.Now(), TRACE, format, args...)
	}
}

// IsTraceEnabled answers if messages on trace level would be logged.
func IsTraceEnabled() bool {
	return TRACE >= threshold
}

// Debug logs message to stderr.
func Debug(s string) {
	if IsDebugEnabled() {
		writeLog(time.Now(), DEBUG, s)
	}
}

// IsDebugEnabled answers if messages on debug level would be logged.
func IsDebugEnabled() bool {
	return DEBUG >= threshold
}

// Debugf logs messages to stderr.
func Debugf(format string, args ...interface{}) {
	if IsDebugEnabled() {
		writeLogf(time.Now(), DEBUG, format, args...)
	}
}

// IsInfoEnabled answers if messages on info level would be logged.
func IsInfoEnabled() bool {
	return INFO >= threshold
}

// Infof logs messages to stderr.
func Infof(format string, args ...interface{}) {
	if IsInfoEnabled() {
		writeLogf(time.Now(), INFO, format, args...)
	}
}

// Info logs messages to stderr.
func Info(s string) {
	if IsInfoEnabled() {
		writeLog(time.Now(), INFO, s)
	}
}

func writeLog(timestamp time.Time, level Level, message string) {
	_, err := fmt.Fprintf(os.Stderr, "%-35s [%-5s] ▶ %s\n", timestamp.Format(time.RFC3339Nano), levelNames[level], message)
	if err != nil {
		panic(err)
	}
}

func writeLogf(timestamp time.Time, level Level, format string, args ...interface{}) {
	writeLog(timestamp, level, fmt.Sprintf(format, args...))
}

func initLogging() Level {
	threshold := getThresholdFromEnv()
	if INFO >= threshold {
		writeLogf(time.Now(), INFO, "Logging threshold: %s", levelNames[threshold])
	}
	return threshold
}
