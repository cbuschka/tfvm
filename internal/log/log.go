package log

import (
	golog "github.com/op/go-logging"
	"os"
)

var logger = initLogger()
var leveledFormattedBackend golog.LeveledBackend

// Trace logs message to stderr.
func Trace(s string) {
	logger.Debug(s)
}

// Tracef logs message to stderr.
func Tracef(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

// IsTraceEnabled answers if messages on trace level would be logged.
func IsTraceEnabled() bool {
	return logger.IsEnabledFor(golog.DEBUG)
}

// Debug logs message to stderr.
func Debug(s string) {
	logger.Info(s)
}

// IsDebugEnabled answers if messages on debug level would be logged.
func IsDebugEnabled() bool {
	return logger.IsEnabledFor(golog.INFO)
}

// Debugf logs messages to stderr.
func Debugf(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

// IsInfoEnabled answers if messages on info level would be logged.
func IsInfoEnabled() bool {
	return logger.IsEnabledFor(golog.NOTICE)
}

// Infof logs messages to stderr.
func Infof(format string, args ...interface{}) {
	logger.Noticef(format, args...)
}

// Info logs messages to stderr.
func Info(args ...interface{}) {
	logger.Notice(args...)
}

// Warningf logs messages to stderr.
func Warningf(format string, args ...interface{}) {
	logger.Warningf(format, args...)
}

// SetVerbosity sets logger log level based on verbosity level.
func SetVerbosity(verbosity int) {
	if verbosity == 0 {
		leveledFormattedBackend.SetLevel(golog.WARNING, "")
	} else if verbosity == 1 {
		leveledFormattedBackend.SetLevel(golog.NOTICE, "")
	} else if verbosity == 2 {
		leveledFormattedBackend.SetLevel(golog.INFO, "")
	} else if verbosity > 2 {
		leveledFormattedBackend.SetLevel(golog.DEBUG, "")
	}
	logger.Noticef("Logging verbosity: %d", verbosity)
}

func initLogger() *golog.Logger {
	format := golog.MustStringFormatter(`%{time:15:04:05.000} [%{level:.4s}] ▶ %{color:reset}%{message}`)
	logger := golog.MustGetLogger("testdrive")
	backend := golog.NewLogBackend(os.Stderr, "", 0)
	formattedBackend := golog.NewBackendFormatter(backend, format)
	leveledFormattedBackend = golog.AddModuleLevel(formattedBackend)
	leveledFormattedBackend.SetLevel(golog.INFO, "")
	logger.SetBackend(leveledFormattedBackend)
	return logger
}
