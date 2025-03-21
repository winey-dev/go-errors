package errors

import (
	"errors"
	"fmt"
	"runtime"

	"github.com/winey-dev/go-errors/codes"
)

type Error struct {
	code    codes.Code
	message string
	err     error
	file    string
	line    int
}

func newError(code codes.Code, message string, err error, caller int) *Error {
	_, file, line, _ := runtime.Caller(caller)
	return &Error{
		code:    code,
		message: message,
		err:     err,
		file:    file,
		line:    line,
	}
}

// New returns an error with the supplied message.
func New(message string) *Error {
	return newError(codes.NotUse, message, nil, 2)
}

func Errorc(code codes.Code, message string) *Error {
	return newError(code, message, nil, 2)
}

// Errorf returns an error with the format and arguments.
func Errorf(format string, args ...any) *Error {
	return newError(codes.NotUse, fmt.Sprintf(format, args...), nil, 2)
}

// Errorcf returns an error with the code and message.
func Errorcf(code codes.Code, format string, args ...any) *Error {
	return newError(code, fmt.Sprintf(format, args...), nil, 2)
}

// Wrap returns an error with the supplied message and the original error.
func Wrap(err error, message string) *Error {
	var Error *Error
	if errors.As(err, &Error) {
		return newError(Error.code, message, err, 2)
	}
	return newError(codes.NotUse, message, err, 2)
}

// Wrapc returns an error with the code and message and the original error.
func Wrapc(err error, code codes.Code, message string) *Error {
	return newError(code, message, err, 2)
}

// Wrapf returns an error with the format and arguments and the original error.
func Wrapf(err error, format string, args ...any) *Error {
	var Error *Error
	if errors.As(err, &Error) {
		return newError(Error.code, fmt.Sprintf(format, args...), err, 2)
	}
	return newError(codes.NotUse, fmt.Sprintf(format, args...), err, 2)
}

func Wrapcf(err error, code codes.Code, format string, args ...any) *Error {
	return newError(code, fmt.Sprintf(format, args...), err, 2)
}

// set the error wrapper
func (e *Error) Wrap(err error) *Error {
	e.err = err
	return e
}

// Wrapcf returns an error with the code, format and arguments and the original error.

// Unwrap returns the wrapper error
func (e *Error) Unwrap() error {
	return e.err
}

// set the error code
func (e *Error) SetCode(code codes.Code) *Error {
	e.code = code
	return e
}

// set the error message
func (e *Error) SetMessage(message string) *Error {
	e.message = message
	return e
}

// Strace returns the error message and the original error message.
func (e *Error) Error() string {
	if e.code == codes.OK {
		return ""
	}

	if e.code == codes.NotUse {
		return e.message
	}
	return fmt.Sprintf("[%s] %s", e.code, e.message)
}

func (e *Error) Details() []string {
	var details []string
	for e.err != nil {
		var Error *Error
		if errors.As(e.err, &Error) {
			details = append(details, Error.Error())
			e = Error
		} else {
			details = append(details, e.err.Error())
			break
		}
	}
	return details
}
