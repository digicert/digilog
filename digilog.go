package digilog

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"strings"
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

// New is used to initialize a new Log
func New() *Log {

	return &Log{
		meta: make(map[string]interface{}),
	}
}

// Log contains loggers for info and error logging as well as the data to be printed in said logs
type Log struct {
	meta map[string]interface{}
}

// AddTag adds metadata to a Log
func (l *Log) AddTag(key string, v interface{}) {
	l.meta[key] = v
}

// AddDuration adds time.Duration to the log metadata
func (l *Log) AddDuration(start time.Time) {
	l.AddTag("duration", fmt.Sprintf("%s", time.Since(start)))
}

// Debug shortcut for log function
func (l *Log) Debug(event string, args ...interface{}) {
	message := fmt.Sprint(args...)
	l.Debugf(event, message)
}

// Debugf shortcut for log function
func (l *Log) Debugf(event string, args ...interface{}) {
	logWriter("DEBUG", l.prepareLog(event, args...))
}

// Info shortcut for log function
func (l *Log) Info(event string, args ...interface{}) {
	message := fmt.Sprint(args...)
	l.Infof(event, message)
}

// Infof shortcut for log function
func (l *Log) Infof(event string, args ...interface{}) {
	logWriter("INFO", l.prepareLog(event, args...))
}

// Warn shortcut for log function
func (l *Log) Warn(event string, args ...interface{}) {
	message := fmt.Sprint(args...)
	l.Warnf(event, message)
}

// Warnf shortcut for log function
func (l *Log) Warnf(event string, args ...interface{}) {
	logWriter("WARN", l.prepareLog(event, args...))
}

// Error shortcut for log function
func (l *Log) Error(event string, args ...interface{}) {
	message := fmt.Sprint(args...)
	l.Errorf(event, message)
}

// Errorf shortcut for log function
func (l *Log) Errorf(event string, args ...interface{}) {
	logWriter("ERROR", l.prepareLog(event, args...))
}

// Fatal is equivalent to calling Error(), then os.Exit(1)
func (l *Log) Fatal(event string, args ...interface{}) {
	l.Errorf(event, args...)
	os.Exit(1)
}

// Critical shortcut for log function
func (l *Log) Critical(event string, args ...interface{}) {
	message := fmt.Sprint(args...)
	l.Criticalf(event, message)
}

// Criticalf shortcut for log function
func (l *Log) Criticalf(event string, args ...interface{}) {
	logWriter("CRITICAL", l.prepareLog(event, args...))

	// Hokey way to test the critical path
	if CriticalExit {
		os.Exit(1)
	}
}

func (l *Log) prepareLog(event string, args ...interface{}) string {
	var message string
	if len(args) > 0 {
		message = fmt.Sprint(args[0])
		if len(args) > 1 {
			message = fmt.Sprintf(message, args[1:len(args)]...)
		}
	}

	logStr := fmt.Sprintf("event_id=%s ", event)

	callStr, err := getCaller()
	if err != nil {
		fmt.Println("error getting caller information", err)
		callStr = ""
	}

	logStr = fmt.Sprintf("%s%s", callStr, logStr)

	for key, val := range l.meta {
		strVal := fmt.Sprintf("%v", val)
		strVal = strings.ReplaceAll(strVal, `"`, "\\\"") // escape double quotes inside strings
		logStr += fmt.Sprintf("%s=%q ", key, strVal)
	}
	l.meta = make(map[string]interface{})

	if len(message) > 0 {
		logStr = fmt.Sprintf("%s%s", logStr, message)
	}

	return logStr
}

func logWriter(loglevel, message string) {
	LogLevelVal := map[string]int{
		"DEBUG":    4,
		"INFO":     3,
		"WARN":     2,
		"ERROR":    1,
		"CRITICAL": 0,
	}

	if LogLevelVal[loglevel] <= LogLevelVal[LogLevel] {
		time := time.Now().Format(time.RFC3339)
		fmt.Fprintf(Out.Out, "%s [%s] %s\n", time, loglevel, message)
	}
}

func getCaller() (string, error) {
	var callStr string
	pc := make([]uintptr, 4)
	numCallers := runtime.Callers(1, pc)
	if numCallers > 0 {
		pc = pc[:numCallers]
		frames := runtime.CallersFrames(pc)
		for {
			frame, more := frames.Next()
			if !more {
				return "", errors.New("digilog: no caller found outside of digilog")
			}
			if !strings.Contains(frame.File, "digilog.go") {
				_, fname := path.Split(frame.File)
				callStr = fmt.Sprintf("%s:%d", fname, frame.Line)
				return callStr, nil
			}
		}
	}

	return "", errors.New("digilog: caller not found")
}
