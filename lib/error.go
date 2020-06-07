package thelm

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"

	"errors"
)

type ErrorHandler struct {
	Out             io.Writer
	PrintStackTrace bool
}

type ThelmError struct {
	Err      error
	Location string
}

func Errorf(format string, a ...interface{}) *ThelmError {
	_, file, line, _ := runtime.Caller(2)
	return &ThelmError{
		Err:      fmt.Errorf(format, a...),
		Location: fmt.Sprintf("%s:%d", file, line),
	}
}

func (e *ThelmError) Error() string {
	return e.Err.Error()
}

// The default error handler
var E = ErrorHandler{
	Out:             os.Stderr,
	PrintStackTrace: true,
}

// New creates a new error variable
func (e *ErrorHandler) New(format string, a ...interface{}) error {
	return Errorf(format, a...)
}

// Annotate increases context information to the error
func (e *ErrorHandler) Annotate(err error, a ...interface{}) error {
	return Errorf("%s: %w", fmt.Sprint(a...), err)
}

// Print writes the error message to predefined io.Writer
func (e *ErrorHandler) Print(err error, a ...interface{}) {
	fmt.Fprint(e.Out, "Error: ")
	if len(a) > 0 {
		fmt.Fprint(e.Out, a...)
		fmt.Fprint(e.Out, ": ")
	}
	fmt.Fprintf(e.Out, "%v\n", err)

	if e.PrintStackTrace {
		sb := strings.Builder{}
		var terr *ThelmError
		for errors.As(err, &terr) {
			sb.WriteString(fmt.Sprintf("  %s\n", terr.Location))
			err = terr
		}
		s := sb.String()
		if s != "" {
			fmt.Fprintf(e.Out, "Error stack:\n%s\n", s)
		}
	}
}

func (e *ErrorHandler) Panic(err error, a ...interface{}) {
	e.Print(err, a...)
	panic("Irrecoverable error")
}
