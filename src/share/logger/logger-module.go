package logger

import (
	log "github.com/sirupsen/logrus"
	"os"
	"sykros-pro/gopro/src/utils"
	"sykros-pro/gopro/src/utils/helper"
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
	logger.context = context
	logger.logger = log.New()
	logger.logger.SetFormatter(&log.JSONFormatter{})
	logger.logger.SetOutput(os.Stdout)

}

func (logger *ViperLogger) commonLogger(data map[string]interface{}) *log.Entry {
	now := time.Now()
	formatted := helper.GetDateTimeFormatter().
		FormatByLayout(now, time.RFC3339)
	defaultObjectMsg := map[string]interface {
	}{
		"time-stamp": formatted,
		"context":    logger.context,
	}
	return logger.logger.WithFields(utils.MergeObject(defaultObjectMsg, data))
}

func (logger *ViperLogger) GetContext() string {
	return logger.context
}

func (logger *ViperLogger) LogWithMsg(msg string, typeLog LogType) {
	now := time.Now()
	formatted := helper.GetDateTimeFormatter().
		FormatByLayout(now, time.RFC3339)
	defaultObjectMsg := map[string]interface {
	}{
		"time-stamp": formatted,
		"context":    logger.context,
	}
	contextLogger := log.WithFields(defaultObjectMsg)
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
