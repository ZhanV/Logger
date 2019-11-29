package logger

import (
	"fmt"
	"os"
)

type ConsoleLogger struct {
	Level int
	file *os.File
}

func NewConsoleLogger()  LogInterface {
	return &ConsoleLogger{
		file: os.Stdout,
	}
}

func (logger *ConsoleLogger) Debug(format string , args ...interface{}) {
	if LogLevelDebug < logger.Level {
		return
	}
	fmt.Fprint(logger.file,WriteLog(LogLevelDebug,format,args...))
}

func (logger *ConsoleLogger) Info(format string , args ...interface{}) {
	if LogLevelInfo < logger.Level {
		return
	}
	fmt.Fprint(logger.file,WriteLog(LogLevelInfo,format,args...))
}

func (logger *ConsoleLogger) Warn(format string , args ...interface{}) {
	if LogLevelWarn < logger.Level {
		return
	}
	fmt.Fprint(logger.file,WriteLog(LogLevelWarn,format,args...))
}

func (logger *ConsoleLogger) Error(format string , args ...interface{}) {
	if LogLevelError < logger.Level {
		return
	}
	fmt.Fprint(logger.file,WriteLog(LogLevelError,format,args...))

}

func (logger *ConsoleLogger) SetLevel(level int)  {
	if level < LogLevelDebug || level > LogLevelError {
		level = LogLevelDebug
	}
	logger.Level = level
}
func (logger *ConsoleLogger) Close()  {
}
