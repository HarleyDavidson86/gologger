package gologger

import "testing"

func Test_Logger(t *testing.T) {
	SetLogLevel(DebugLevel)
	message := "This is a test"
	// "PANIC"
	Panic(message)
	// "ERROR"
	Error(message)
	// "WARN"
	Warn(message)
	// "INFO"
	Info(message)
	// "DEBUG"
	Debug(message)
	// "TRACE"
	Trace(message)
}

func Test_LoggerFmt(t *testing.T) {
	SetLogLevel(DebugLevel)
	message := "This is a %s"
	// "PANIC"
	Panicf(message, "Panic Message")
	// "ERROR"
	Errorf(message, "Error Message")
	// "WARN"
	Warnf(message, "Warn Message")
	// "INFO"
	Infof(message, "Info Message")
	// "DEBUG"
	Debugf(message, "Debug Message")
	// "TRACE"
	Tracef(message, "Trace Message")
}
