package common

import (
	"errors"
	"net/http"
)

type Errkey int64

const (
	DB_ERR          Errkey = 0
	INTERNAL               = 1
	INVALID_REQUEST        = 2
)

var (
	DB_err = errors.New("DB_ERR")
)

func getErrKey(key Errkey) string {
	return [...]string{
		"DB_ERR", "INTERNAL_SERVER_ERROR", "INVALID_REQUEST",
	}[key]
}

type AppError struct {
	StatusCode int    `json:"status_code"`
	RootError  error  `json:"_"`
	Log        string `json:"log"`
	Key        string `json:"key"`
	Message    string `json:"message"`
}

func NewErrorResponse(msg, log, key string, rootErr error) *AppError {
	return &AppError{
		http.StatusBadRequest,
		rootErr,
		log,
		key,
		msg,
	}
}

func FullErrResponse(msg, log, key string, rootErr error, statusCode int) *AppError {
	return &AppError{
		statusCode,
		rootErr,
		log,
		key,
		msg,
	}
}

func (appError *AppError) RootErr() error {
	if err, ok := appError.RootError.(*AppError); ok {
		return err.RootErr()
	}
	return appError.RootError
}

func (appError *AppError) Error() string {
	return appError.RootError.Error()
}
func ErrInvalidRequest(e error) *AppError {
	return NewErrorResponse("INVALID REQUEST",
		e.Error(),
		getErrKey(INVALID_REQUEST),
		e)
}

func ErrInternalServerError(error error) *AppError {
	return FullErrResponse("Internal server error",
		error.Error(),
		getErrKey(INTERNAL),
		error,
		http.StatusInternalServerError)
}

func ErrDB(e error) *AppError {
	return NewErrorResponse("SOMETHING WRONG WITH DB",
		e.Error(),
		getErrKey(DB_ERR),
		e)

}
