package digilog

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"time"
)

// LogLevel sets the log level to log
var LogLevel string

// CriticalExit makes Critical funcs exit when calling
var CriticalExit bool

// Out prints the data to os.Stdout/os.StdErr
var Out *BuffOut

// BuffOut provides writers to handle output and err output
type BuffOut struct {
	Out io.Writer
	Err io.Writer
}

func init() {
	LogLevel = os.Getenv("LOG_LEVEL")
	if LogLevel == "" {
		LogLevel = "DEBUG"
	}

	CriticalExit = true

	Out = &BuffOut{Out: os.Stdout, Err: os.Stderr}
}

// Debug shortcut for log function
func Debug(args ...interface{}) {
	message := fmt.Sprint(args...)
	Debugf(message)
}

// Debugf shortcut for log function
func Debugf(message string, args ...interface{}) {
	logIt("DEBUG", message, args...)
}

// Info shortcut for log function
func Info(args ...interface{}) {
	message := fmt.Sprint(args...)
	Infof(message)
}

// Infof shortcut for log function
func Infof(message string, args ...interface{}) {
	logIt("INFO", message, args...)
}

// Warn shortcut for log function
func Warn(args ...interface{}) {
	message := fmt.Sprint(args...)
	Warnf(message)
}

// Warnf shortcut for log function
func Warnf(message string, args ...interface{}) {
	logIt("WARN", message, args...)
}

// Error shortcut for log function
func Error(args ...interface{}) {
	message := fmt.Sprint(args...)
	Errorf(message)
}

// Errorf shortcut for log function
func Errorf(message string, args ...interface{}) {
	logIt("ERROR", message, args...)
}

// Critical shortcut for log function
func Critical(args ...interface{}) {
	message := fmt.Sprint(args...)
	Criticalf(message)
}

// Criticalf shortcut for log function
func Criticalf(message string, args ...interface{}) {
	logIt("CRITICAL", message, args...)

	// Hokey way to test the critical path
	if CriticalExit {
		os.Exit(1)
	}
}

// lame logger
func logIt(loglevel string, message string, args ...interface{}) {
	LogLevelVal := map[string]int{
		"DEBUG":    4,
		"INFO":     3,
		"WARN":     2,
		"ERROR":    1,
		"CRITICAL": 0,
	}

	_, file, line, _ := runtime.Caller(2)

	if len(args) > 0 {
		message = fmt.Sprintf(message, args...)
	}

	if LogLevelVal[loglevel] <= LogLevelVal[LogLevel] {
		time := time.Now().Format(time.RFC3339)
		fmt.Fprintf(Out.Out, "%s file=%s line=%d %s: %s\n", time, file, line, loglevel, message)
	}
}
