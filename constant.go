package logger

const (
	LogLevelDebug = iota
	LogLevelInfo
	LogLevelWarn
	LogLevelError
)

const (
	LogSplitHour = "HOUR"
	LogSplitSize = "SIZE"

)


func getLogLevelName(level int) string {
	switch level {
	case LogLevelDebug:
		return "DEBUG"
	case LogLevelInfo:
		return "INFO"
	case LogLevelWarn:
		return "WARN"
	case LogLevelError:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}
func getLogLevelType(level string) int {
	switch level {
	case  "DEBUG":
		return LogLevelDebug
	case  "INFO":
		return LogLevelInfo
	case  "WARN":
		return LogLevelWarn
	case  "ERROR":
		return LogLevelError
	default:
		return 0
	}
}