package errors

import (
	"errors"
	"fmt"
)

type Error struct {
	ErrorLevel string
	ErrorType  string
	Info       error
}

const (
	ServerErrorLevel   = "server_error"
	DataBaseErrorLevel = "database_error"
	ApiErrorLevel      = "api_error"
)

func SetError(errorLevel, errorType string, info error) error {
	err := &Error{
		ErrorLevel: errorLevel,
		ErrorType:  errorType,
		Info:       info,
	}
	return err.convertError()
}

func (err *Error) convertError() error {
	if err.Info != nil {
		return errors.New(fmt.Sprintf("%s.%s (%s)", err.ErrorLevel, err.ErrorType, err.Info))
	}
	return errors.New(fmt.Sprintf("%s.%s", err.ErrorLevel, err.ErrorType))
}
