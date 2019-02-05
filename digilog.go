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

	Out = &BuffOut{Out: os.Stdout, Err: os.Stderr}
}

// Debug shortcut for log function
func Debug(message string, args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = ""
		line = 0
	}
	log("DEBUG", file, line, message, args...)
}

// Info shortcut for log function
func Info(message string, args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = ""
		line = 0
	}
	log("INFO", file, line, message, args...)
}

// Warn shortcut for log function
func Warn(message string, args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = ""
		line = 0
	}
	log("WARN", file, line, message, args...)
}

// Error shortcut for log function
func Error(message string, args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = ""
		line = 0
	}
	log("ERROR", file, line, message, args...)
}

// Critical shortcut for log function
func Critical(message string, args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = ""
		line = 0
	}
	log("CRITICAL", file, line, message, args...)
}

// Log lame Log function
func log(loglevel string, file string, line int, message string, args ...interface{}) {
	LogLevelVal := map[string]int{
		"DEBUG":    4,
		"INFO":     3,
		"WARN":     2,
		"ERROR":    1,
		"CRITICAL": 0,
	}

	if LogLevelVal[loglevel] <= LogLevelVal[LogLevel] {
		time := time.Now().Format(time.RFC3339)
		fmt.Fprintf(Out.Out, "%s file=%s line=%d %s: %s\n", time, file, line, loglevel, fmt.Sprintf(message, args...))
	}

	if LogLevelVal[loglevel] == 0 {
		os.Exit(1)
	}
}
