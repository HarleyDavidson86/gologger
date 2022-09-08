package gologger

import (
	"fmt"
	"os"
	"time"
)

func log(level Level, logargs ...interface{}) {
	if IsLevelEnabled(level) {
		//2006-01-02 15:04:05.000 [DEBUG  ] - Logargs
		fmt.Printf("%s [%-7s] - %s\n",
			time.Now().Format(timestampFormat),
			ParseLogLevelToString(level),
			fmt.Sprint(logargs...))
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

// Prints out Fatal args and terminates application
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

func Trace(args ...interface{}) {
	log(TraceLevel, args...)
}

func Debug(args ...interface{}) {
	log(DebugLevel, args...)
}

func Info(args ...interface{}) {
	log(InfoLevel, args...)
}

func Warn(args ...interface{}) {
	log(WarnLevel, args...)
}

func Warning(args ...interface{}) {
	Warn(args...)
}

func Error(args ...interface{}) {
	log(ErrorLevel, args...)
}

// Prints out Fatal args and terminates application
func Fatal(args ...interface{}) {
	log(FatalLevel, args...)
	os.Exit(1)
}

func Panic(args ...interface{}) {
	log(PanicLevel, args...)
}
