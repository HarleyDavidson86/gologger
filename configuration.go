package gologger

import stdlog "log"

var loglevel Level = DebugLevel

var timestampFormat = "2006-01-02 15:04:05.000"

// Set the loglevel
func SetLogLevel(level Level) {
	loglevel = level
}

// Set the timestamp format (see time.Format)
// and sets flags of log package to 0
func SetTimestampFormat(format string) {
	stdlog.SetFlags(0)
	timestampFormat = format
}
