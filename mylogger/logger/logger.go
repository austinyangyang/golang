package logger

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
	"time"
)

// log level custom
type LogLevel uint16

//log level const
const (
	UNKNOWN LogLevel = iota
	DEBUG
	TRACE
	INFO
	WARRING
	ERROR
	FATAL
)

//Logger 日志类
type Logger struct {
	Level LogLevel
}

//NewLogger 构造函数
func NewLog(levelStr string) Logger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)

	}

	return Logger{
		Level: level,
	}
}

func parseLogLevel(s string) (LogLevel, error) {
	s = strings.ToLower(s)

	switch s {
	case "debug":
		return DEBUG, nil
	case "trace":
		return TRACE, nil
	case "info":
		return INFO, nil
	case "warring":
		return WARRING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		err := errors.New("日志级别错误")
		return UNKNOWN, err
	}

}

func getLogString(lv LogLevel) string {
	switch lv {
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	case INFO:
		return "INFO"
	case WARRING:
		return "WARRING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "UNKNOWN"

	}
}

func (l Logger) enabled(logLevel LogLevel) bool {
	return logLevel >= l.Level
}




func (l *Logger) logPrint(lv LogLevel, format string, args ...interface{}) {
	var times string
	if l.enabled(lv) {

		msg := fmt.Sprintf(format, args...)
		now := time.Now()
		times = now.Format("2006-01-02 15:04:05")
		funcName, fileName, lineNumber := getInfo(3)

		fmt.Printf("[%s] [%s] [%s %s %d] %s\n", times, getLogString(lv), funcName, fileName, lineNumber, msg)
	}
}

func (l Logger) Trace(format string, a ...interface{}) {

	l.logPrint(TRACE, format, a...)

}

func (l Logger) Debug(format string, a ...interface{}) {

	l.logPrint(DEBUG, format, a...)

}

func (l Logger) Info(format string, a ...interface{}) {

	l.logPrint(INFO, format, a...)

}

func (l Logger) Warring(format string, a ...interface{}) {

	l.logPrint(WARRING, format, a...)

}

func (l Logger) Error(format string, a ...interface{}) {

	l.logPrint(ERROR, format, a...)

}

func (l Logger) Fatal(format string, a ...interface{}) {

	l.logPrint(FATAL, format, a...)

}

func getInfo(skip int) (funcName, fileName string, lineNumber int) {

	pc, file, lineNumber, ok := runtime.Caller(skip)
	if !ok {
		fmt.Printf("runtime.Caller() failed \n")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	funcName = strings.Split(funcName, ".")[0]
	fileName = path.Base(file)

	return funcName, fileName, lineNumber
}
