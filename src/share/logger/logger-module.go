package logger

import (
	"github.com/mattn/go-colorable"
	log "github.com/sirupsen/logrus"
	"github.com/ttys3/rotatefilehook"
	_ "os"
	"sykros-pro/gopro/src/utils"
	"time"
)

type LogType int64

const (
	INFO  LogType = 0
	WARN          = 1
	ERROR         = 2
)

type LoggerService interface {
	InitLogger(context string)
	GetContext() string
	commonLogger(data map[string]interface{}) *log.Entry
	Info(data map[string]interface{})
	Warn(data map[string]interface{})
	Error(data map[string]interface{})
	LogWithMsg(msg string, typeLog LogType)
}

type ViperLogger struct {
	LoggerService
	context string
	logger  *log.Logger
}

func (logger *ViperLogger) InitLogger(context string) {
	rotateFileHook, err := rotatefilehook.NewRotateFileHook(rotatefilehook.RotateFileConfig{
		Filename:   "logfile.log",
		MaxSize:    5,
		MaxBackups: 7,
		MaxAge:     7,
		LocalTime:  true,
		Level:      log.InfoLevel,
		Formatter:  &log.TextFormatter{FullTimestamp: true},
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	logger.context = context
	logger.logger = log.New()
	logger.logger.SetFormatter(&log.TextFormatter{
		ForceColors:     true,
		DisableColors:   false,
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339,
		FieldMap: log.FieldMap{
			"FieldKeyTime":  "@timestamp",
			"FieldKeyLevel": "@level",
			"FieldKeyMsg":   "@message",
			"FieldKeyFunc":  "@caller",
		},
	})
	logger.logger.AddHook(rotateFileHook)
	logger.logger.SetOutput(colorable.NewColorableStdout())

}

func (logger *ViperLogger) commonLogger(data map[string]interface{}) *log.Entry {
	defaultObjectMsg := map[string]interface {
	}{
		"context": logger.context,
	}
	return logger.logger.WithFields(utils.MergeObject(defaultObjectMsg, data))
}

func (logger *ViperLogger) GetContext() string {
	return logger.context
}

func (logger *ViperLogger) LogWithMsg(msg string, typeLog LogType) {
	contextLogger := logger.logger.WithFields(map[string]interface{}{})
	switch typeLog {
	case INFO:
		contextLogger.Info(msg)
	case WARN:
		contextLogger.Warn(msg)
	case ERROR:
		contextLogger.Error(msg)

	}
}

func (logger *ViperLogger) Info(data map[string]interface{}) {
	logger.commonLogger(data).Info()
}

func (logger *ViperLogger) Warn(data map[string]interface{}) {
	logger.commonLogger(data).Info()
}

func (logger *ViperLogger) Error(data map[string]interface{}) {
	logger.commonLogger(data).Info()
}

func LogrusSetup(context string) LoggerService {
	viperLogger := &ViperLogger{}
	viperLogger.logger = log.New()
	viperLogger.InitLogger(context)
	return viperLogger
}

//type TestStruct struct {
//	Index int
//	Flag bool
//}
//
//func Test(v any){
//	switch  v.(type) {
//	case *TestStruct:
//		if v.(*TestStruct).Flag == false {
//			v.(*TestStruct).Flag = true
//			v.(*TestStruct).Index = 10
//		}
//	}
//}
