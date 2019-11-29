package logger

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

func GetLineInfo()(fileName string, funcName string , lineNo int)  {
	pc,file,line,ok := runtime.Caller(4)
	if ok {
		fileName = file
		funcName = runtime.FuncForPC(pc).Name()
		lineNo = line
	}
	return
}



func WriteLog(level int, format string , args ...interface{}) string  {


	timeStamp := time.Now().Format("2006-01-02 15:04:05.999")
	fileName, funcName, lineNo := GetLineInfo()
	fileName = filepath.Base(fileName)
	funcName = filepath.Base(funcName)

	content := fmt.Sprintf(format,args...)
	content = fmt.Sprintf("[%s][%s] [%s:%s:%d] : %s \r\n",getLogLevelName(level), timeStamp,fileName, funcName, lineNo, content)

	return content
}



//读取key=value类型的配置文件
func InitConfig(path string) map[string]string {
	config := make(map[string]string)

	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		panic(err)
	}

	r := bufio.NewReader(f)
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		s := strings.TrimSpace(string(b))
		index := strings.Index(s, "=")
		if index < 0 {
			continue
		}
		key := strings.TrimSpace(s[:index])
		if len(key) == 0 {
			continue
		}
		value := strings.TrimSpace(s[index+1:])
		if len(value) == 0 {
			continue
		}
		config[key] = value
	}
	return config
}
