package logger

import (
	"log"
	"os"
)

type LoggerService interface {
	initLogger()
	error(msg string)
	warn(msg string)
	info(msg string)
	GetContext() string
}

type ViperLogger struct {
	LoggerService
	infoLog  *log.Logger
	errorLog *log.Logger
	warnLog  *log.Logger
	context  string
}

func (l *ViperLogger) error(msg string) {
	l.errorLog.Printf("[%s]%s\n", l.context, msg)
}
func (l *ViperLogger) warn(msg string) {
	l.warnLog.Printf("[%s]%s\n", l.context, msg)
}
func (l *ViperLogger) info(msg string) {
	l.infoLog.Printf("[%s]%s\n", l.context, msg)
}

func (l *ViperLogger) GetContext() string {
	return l.context
}

func InitLogger(context string) LoggerService {
	loggerInstance := &ViperLogger{}
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	loggerInstance.context = context
	loggerInstance.infoLog = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	loggerInstance.warnLog = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	loggerInstance.errorLog = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	return loggerInstance
}
