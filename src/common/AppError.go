package common

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
)

type Errkey int64

const (
	DB_ERR             Errkey = 0
	INTERNAL                  = 1
	INVALID_REQUEST           = 2
	INVALID_PARAM             = 3
	CAN_NOT_GET_ENTITY        = 4
)

var (
	DB_err = errors.New("DB_ERR")
)

func getErrKey(key Errkey) string {
	return [...]string{
		"DB_ERR", "INTERNAL_SERVER_ERROR", "INVALID_REQUEST", "INVALID_PARAM", "CAN_NOT_GET_ENTITY",
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

func InvalidTypeOfParam(param interface{}, expectType reflect.Kind, err error) *AppError {
	msg := fmt.Sprintf("Invalid Param %v of type %v", param, reflect.ValueOf(param), expectType)
	return NewErrorResponse(
		msg,
		err.Error(),
		getErrKey(INVALID_PARAM),
		err,
	)
}

func CanNotGetEntity(entity string, err error) *AppError {
	fmt.Printf("Can not get %s\n", entity)
	return NewErrorResponse(
		"Can not get "+entity,
		err.Error(),
		getErrKey(CAN_NOT_GET_ENTITY),
		err,
	)
}
