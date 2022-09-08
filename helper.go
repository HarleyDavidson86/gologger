package gologger

func IsLevelEnabled(level Level) bool {
	return loglevel >= level
}

func ParseLogLevelToString(level Level) string {
	switch level {
	case PanicLevel:
		return "PANIC"
	case FatalLevel:
		return "FATAL"
	case ErrorLevel:
		return "ERROR"
	case WarnLevel:
		return "WARN"
	case InfoLevel:
		return "INFO"
	case DebugLevel:
		return "DEBUG"
	default:
		return "TRACE"
	}
}
