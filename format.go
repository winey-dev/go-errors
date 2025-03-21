package errors

import (
	"errors"
	"fmt"
	"strings"

	"github.com/winey-dev/go-errors/codes"
)

func (e *Error) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') && s.Flag('#') {
			fmt.Fprint(s, e.RuntimeTrace())
			return
		}
		if s.Flag('+') {
			fmt.Fprint(s, e.Trace())
			return
		}
		if s.Flag('#') {
			fmt.Fprint(s, e.FileLineTrace())
			return
		}
		fallthrough
	case 's':
		fmt.Fprint(s, e.Error())
	case 'q':
		fmt.Fprintf(s, "%q", e.Error())
	}
}

// FileLineTrace returns the error message with file and line information.
func (e *Error) FileLineTrace() string {
	if e.code == codes.NotUse {
		return fmt.Sprintf("%s (%s:%d)", e.message, e.file, e.line)
	}
	return fmt.Sprintf("[%s] %s (%s:%d)", e.code, e.message, e.file, e.line)
}

// RuntimeTrace returns the error message with runtime information.
func (e *Error) RuntimeTrace() string {
	var trace strings.Builder
	if e.code == codes.NotUse {
		trace.WriteString(fmt.Sprintf("%s (%s:%d)", e.message, e.file, e.line))
	} else {
		trace.WriteString(fmt.Sprintf("[%s] %s (%s:%d)", e.code, e.message, e.file, e.line))
	}

	for e.err != nil {
		var wrappedErr *Error
		if errors.As(e.err, &wrappedErr) {
			trace.WriteString(fmt.Sprintf("\n\t-> %s (%s:%d)", wrappedErr, wrappedErr.file, wrappedErr.line))
			e = wrappedErr
		} else {
			trace.WriteString(fmt.Sprintf("\n\t-> %v", e.err))
			break
		}
	}

	return trace.String()
}

// Trace returns the error message and the original error message.
func (e *Error) Trace() string {
	var trace []string
	if e.code == codes.NotUse {
		trace = append(trace, e.message)
	} else {
		trace = append(trace, fmt.Sprintf("[%s] %s", e.code, e.message))
	}

	for e.err != nil {
		var wrappedErr *Error
		if errors.As(e.err, &wrappedErr) {
			trace = append(trace, fmt.Sprintf("\t-> %s", wrappedErr))
			e = wrappedErr
		} else {
			trace = append(trace, fmt.Sprintf("\t-> %v", e.err))
			break
		}
	}

	return strings.Join(trace, "\n")
}
