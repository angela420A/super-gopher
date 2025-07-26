package logger

import (
	"log"
	"os"
)

var (
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
)

func init() {
	infoLogger = log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime)
	warnLogger = log.New(os.Stderr, "[WARN] ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(os.Stdout, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)
}

func Info(msg string, v ...interface{}) {
	infoLogger.Println(msg)
}

func Warn(msg string, v ...interface{}) {
	warnLogger.Println(msg)
}

func Error(msg string, err error, v ...interface{}) {
	if err != nil {
		errorLogger.Printf("%s: %v", msg, err)
	} else {
		errorLogger.Print(msg)
	}
}
