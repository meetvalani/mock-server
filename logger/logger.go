package mockserver

import (
	"log"
	"os"
)

var (
	warningLogger *log.Logger
	infoLogger    *log.Logger
	errorLogger   *log.Logger
)

func Info(format string, values ...interface{}) {
	infoLogger.Printf(format, values...)
}
func Warning(format string, values ...interface{}) {
	warningLogger.Printf(format, values...)
}
func Error(format string, values ...interface{}) {
	errorLogger.Printf(format, values...)
}

func SetUp(fileName string) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	infoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
