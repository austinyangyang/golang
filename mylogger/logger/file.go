package logger

import (
	"fmt"
	"os"
	"path"
	"time"
)

type FileLogger struct {
	Level       LogLevel
	fileName    string
	filePath    string
	fileObj     *os.File
	errfileObj  *os.File
	maxFileSize int64
}

func NewFileLogger(levelStr, fp, fn string, maxSize int64) *FileLogger {
	logLevel, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	fl := &FileLogger{
		Level:       logLevel,
		fileName:    fn,
		filePath:    fp,
		maxFileSize: maxSize,
	}
	err = fl.initFile()

	if err != nil {
		panic(err)
	}
	return fl

}

func (fl *FileLogger) checkFileSize(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed, error: %v\n", err)
		return false
	}
	return fileInfo.Size() >= fl.maxFileSize

}
func (fl *FileLogger) initFile() error {
	fullFileName := path.Join(fl.filePath, fl.fileName)

	fileobj, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		fmt.Printf("open log file failed %v \n", err)
		return err
	}

	errfileobj, err := os.OpenFile(fullFileName+".error", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		fmt.Printf("open err log file failed %v \n", err)
		return err
	}
	fl.fileObj = fileobj
	fl.errfileObj = errfileobj
	return nil

}

func (fl *FileLogger) FileSplit(file *os.File) (*os.File, error) {

	nowstr := time.Now().Format("200601021504")

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed, error: %v\n", err)
		return nil, err
	}

	logName := path.Join(fl.filePath, fileInfo.Name())
	newLogName := fmt.Sprintf("%s.bak%s", logName, nowstr)
	file.Close()

	os.Rename(logName, newLogName)

	fileobj, err := os.OpenFile(newLogName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		fmt.Printf("open err log file failed %v \n", err)
		return nil, err
	}

	return fileobj, nil

}

func (fl *FileLogger) logPrint(lv LogLevel, format string, args ...interface{}) {

	if fl.enabled(lv) {

		var times string
		msg := fmt.Sprintf(format, args...)
		now := time.Now()
		times = now.Format("2006-01-02 15:04:05")
		funcName, fileName, lineNumber := getInfo(3)

		if fl.checkFileSize(fl.fileObj) {
			newFile, err := fl.FileSplit(fl.fileObj)
			if err != nil {
				return
			}
			fl.fileObj = newFile

		}

		fmt.Fprintf(fl.fileObj, "[%s] [%s] [%s %s %d] %s\n", times, getLogString(lv), funcName, fileName, lineNumber, msg)
		if lv >= ERROR {

			if fl.checkFileSize(fl.errfileObj) {
				newFile, err := fl.FileSplit(fl.errfileObj)
				if err != nil {
					return
				}
				fl.fileObj = newFile

			}
			fmt.Fprintf(fl.errfileObj, "[%s] [%s] [%s %s %d] %s\n", times, getLogString(lv), funcName, fileName, lineNumber, msg)

		}
	}
}

func (fl *FileLogger) enabled(logLevel LogLevel) bool {
	return logLevel >= fl.Level
}

func (fl FileLogger) Trace(format string, a ...interface{}) {

	fl.logPrint(TRACE, format, a...)

}

func (fl FileLogger) Debug(format string, a ...interface{}) {

	fl.logPrint(DEBUG, format, a...)

}

func (fl FileLogger) Info(format string, a ...interface{}) {

	fl.logPrint(INFO, format, a...)

}

func (fl FileLogger) Warring(format string, a ...interface{}) {

	fl.logPrint(WARRING, format, a...)

}

func (fl FileLogger) Error(format string, a ...interface{}) {

	fl.logPrint(ERROR, format, a...)

}

func (fl FileLogger) Fatal(format string, a ...interface{}) {

	fl.logPrint(FATAL, format, a...)

}

func (fl *FileLogger) Close() {
	fl.fileObj.Close()
	fl.errfileObj.Close()
}
