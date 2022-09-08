package gologger

var loglevel Level = DebugLevel
var timestampFormat = "2006-01-02 15:04:05.000"

// Set the loglevel
func SetLogLevel(level Level) {
	loglevel = level
}

// Set the timestamp format (see time.Format)
func SetTimestampFormat(format string) {
	timestampFormat = format
}
