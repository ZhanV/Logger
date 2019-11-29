package logger

import (
	"fmt"
	"os"
	"time"
)

type FileLogger struct {
	Level int
	FilePath string
	FileName string
	file *os.File
	logChan chan string
	logSplitType string
	logSplitSize int64
	lastLogHour int
}

func NewFileLogger(path string, name string, logChanSize int, logSplitType string ,logSplitSize int64)  LogInterface {
	var fileName = fmt.Sprintf("%s\\%s.log",path,name)
	logFile ,err := os.OpenFile(fileName,os.O_CREATE | os.O_APPEND| os.O_WRONLY,0755)
	if err != nil {
		 panic(fmt.Sprintf("open file [%s] failure : %s ", fileName,err))
	}

	logger := &FileLogger{
		FilePath: path,
		FileName: name,
		file: logFile,
		logChan: make(chan string, logChanSize),
		logSplitType:logSplitType,
		logSplitSize:logSplitSize,
		lastLogHour: time.Now().Hour(),
	}



	go logger.WriteLogToFile()

	return logger
}


func (logger *FileLogger) Debug(format string , args ...interface{}) {
	if LogLevelDebug < logger.Level {
		return
	}

	var content = WriteLog(LogLevelDebug,format,args...)
	select {
	case logger.logChan <- content:
	default:
	}
}

func (logger *FileLogger) Info(format string , args ...interface{}) {
	if LogLevelInfo < logger.Level {
		return
	}
	var content = WriteLog(LogLevelInfo,format,args...)
	select {
	case logger.logChan <- content:
	default:
	}
}

func (logger *FileLogger) Warn(format string , args ...interface{}) {
	if LogLevelWarn < logger.Level {
		return
	}
	var content = WriteLog(LogLevelWarn,format,args...)
	select {
	case logger.logChan <- content:
	default:
	}
}

func (logger *FileLogger) Error(format string , args ...interface{}) {
	if LogLevelError < logger.Level {
		return
	}
	var content = WriteLog(LogLevelError,format,args...)
	select {
	case logger.logChan <- content:
	default:
	}
}

func (logger *FileLogger) SetLevel(level int)  {
	if level < LogLevelDebug || level > LogLevelError {
		level = LogLevelDebug
	}
	logger.Level = level
}

func (logger *FileLogger) Close()  {
	logger.file.Close()
}

func (logger *FileLogger) WriteLogToFile() {
	for logContent := range logger.logChan {

		logger.SplitLogFile();

		fmt.Fprint(logger.file, logContent)
	}
}

func (logger *FileLogger) SplitLogFile() {

	if logger.logSplitType == LogSplitHour {
		logger.splitLogFileHour()
	} else {
		logger.SplitLogFileSize()
	}

}

func (logger *FileLogger) splitLogFileHour() {
	var backupFileName string
	var fileName string
	now := time.Now()
	hour := now.Hour()
	if hour == logger.lastLogHour {
		return
	}
	backupFileName = fmt.Sprintf("%s\\%s.log_%04d%02d%02d%02d", logger.FilePath, logger.FileName, now.Year(), now.Month(), now.Day(), logger.lastLogHour)
	fileName = fmt.Sprintf("%s\\%s.log", logger.FilePath, logger.FileName)
	logger.file.Close()
	os.Rename(fileName, backupFileName)
	logFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		return
	}
	logger.file = logFile
	logger.lastLogHour = hour
}

func (logger *FileLogger) SplitLogFileSize() {
	var backupFileName string
	var fileName string
	now := time.Now()
	stat, err := logger.file.Stat()
	if err != nil {
		return
	}

	if logger.logSplitSize > stat.Size() {
		return
	}

	backupFileName = fmt.Sprintf("%s\\%s.log_%04d%02d%02d%02d%02d%02d", logger.FilePath, logger.FileName, now.Year(), now.Month(), now.Day(), now.Hour(),now.Minute(),now.Second())
	fileName = fmt.Sprintf("%s\\%s.log", logger.FilePath, logger.FileName)
	logger.file.Close()
	os.Rename(fileName, backupFileName)
	logFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		return
	}
	logger.file = logFile
}
