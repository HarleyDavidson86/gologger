package gologger

import (
	"fmt"
	"os"
	"time"
)

func log(level Level, logmessage string) {
	if IsLevelEnabled(level) {
		//2006-01-02 15:04:05.000 [DEBUG  ] - Logmessage
		fmt.Printf("%s [%-7s] - %s\n",
			time.Now().Format(timestampFormat),
			ParseLogLevelToString(level),
			logmessage)
	}
}

func Tracef(format string, args ...interface{}) {
	logf(TraceLevel, format, args...)
}

func Debugf(format string, args ...interface{}) {
	logf(DebugLevel, format, args...)
}

func Infof(format string, args ...interface{}) {
	logf(InfoLevel, format, args...)
}

func Warnf(format string, args ...interface{}) {
	logf(WarnLevel, format, args...)
}

func Warningf(format string, args ...interface{}) {
	Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	logf(ErrorLevel, format, args...)
}

// Prints out Fatal message and terminates application
func Fatalf(format string, args ...interface{}) {
	logf(FatalLevel, format, args...)
	os.Exit(1)
}

func Panicf(format string, args ...interface{}) {
	logf(PanicLevel, format, args...)
}

func logf(level Level, format string, args ...interface{}) {
	log(level, fmt.Sprintf(format, args...))
}

func Trace(message string) {
	log(TraceLevel, message)
}

func Debug(message string) {
	log(DebugLevel, message)
}

func Info(message string) {
	log(InfoLevel, message)
}

func Warn(message string) {
	log(WarnLevel, message)
}

func Warning(message string) {
	Warn(message)
}

func Error(message string) {
	log(ErrorLevel, message)
}

// Prints out Fatal message and terminates application
func Fatal(message string) {
	log(FatalLevel, message)
	os.Exit(1)
}

func Panic(message string) {
	log(PanicLevel, message)
}
