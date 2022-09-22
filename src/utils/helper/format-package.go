package helper

import "time"

type Provider int64

const (
	TIME_FORMATTER   Provider = 0
	LOGGER_FORMATTER          = 1
)

type BaseProvider interface {
	getIndex() Provider
}

type baseProvider struct {
	index Provider
}

func (b *baseProvider) getIndex() Provider {
	return b.index
}

type DateTimeFormatter interface {
	BaseProvider
	DefaultFormat(inputTime time.Time) string
	FormatByLayout(inputTime time.Time, format string) string
}

type dateTimeFormatter struct {
	DateTimeFormatter
	*baseProvider
}

func (t *dateTimeFormatter) FormatByLayout(inputTime time.Time, format string) string {
	return inputTime.Format(format)
}

func (t *dateTimeFormatter) DefaultFormat(inputTime time.Time) string {
	return inputTime.Format(time.RFC3339)
}


func (t *dateTimeFormatter) getIndex() Provider {
	return t.baseProvider.getIndex()
}

func GetDateTimeFormatter() DateTimeFormatter {
	return &dateTimeFormatter{
		baseProvider: &baseProvider{
			index: TIME_FORMATTER,
		},
	}
}
