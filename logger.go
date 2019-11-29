package logger

import (
	"fmt"
	"strconv"
)

var log LogInterface

func Debug(format string, args ...interface{})  {
	log.Debug(format , args...);
}

func Info(format string, args ...interface{})  {
	log.Info(format , args...);
}

func Warn(format string, args ...interface{})  {
	log.Warn(format , args...);
}

func Error(format string, args ...interface{})  {
	log.Error(format , args...);
}

func Init() (err error) {
	config := InitConfig("conf.ini")
	switch config["name"] {
	case "file":

		var logSplitType = LogSplitHour
		var logSplitSize int64
		if _,ok := config["file_path"]; !ok{
			fmt.Errorf("can't get param file_path")
			return
		}
		if _,ok := config["file_name"]; !ok{
			fmt.Errorf("can't get param file_name")
			return
		}
		if _,ok := config["log_chan_size"]; !ok{
			config["log_chan_size"] = "50000" //通道默认容量50000
		}

		if _,ok := config["log_split_type"];ok{
			if config["log_split_type"] == LogSplitSize {
				logSplitType = LogSplitSize

				if _, ok := config["log_split_size"] ; !ok {
					// 如果没有设置切割大小，默认为100M
					logSplitSize = 104857600
				} else {
					logSplitSize , err = strconv.ParseInt(config["log_split_size"],10, 64)
					if	err != nil {
						fmt.Errorf(" set param log_split_size error : %s", err)
						return
					}
				}
			}
		}

		var logChanSize,_ = strconv.Atoi(config["log_chan_size"])
		log =  NewFileLogger(config["file_path"], config["file_name"],logChanSize, logSplitType, logSplitSize )
	case "console":
		log = NewConsoleLogger()
	default:
		err = fmt.Errorf("unknown log type %s", config["name"])
		return
	}

	log.SetLevel(getLogLevelType(config["log_level"]));
	return
}



