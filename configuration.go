package gologger

var loglevel Level = DebugLevel

// Set the loglevel
func SetLogLevel(level Level) {
	loglevel = level
}
