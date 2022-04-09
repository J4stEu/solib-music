package errors

import (
	"errors"
	"fmt"
)

type Error struct {
	ErrorLevel string
	ErrorType  string
}

const (
	ServerErrorLevel   = "server_error"
	DataBaseErrorLevel = "database_error"
	ApiErrorLevel      = "api_error"
)

func SetError(errorLevel, errorType string) error {
	err := &Error{
		ErrorLevel: errorLevel,
		ErrorType:  errorType,
	}
	return err.convertError()
}

func (err *Error) convertError() error {
	return errors.New(fmt.Sprintf("%s.%s", err.ErrorLevel, err.ErrorType))
}
