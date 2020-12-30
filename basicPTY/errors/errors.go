package errors

import "fmt"

type Exit struct {
	msg  string
	err  error
	Code int
}

func NewExitError(code int) *Exit {
	return &Exit{Code: code}
}

func WrapExitError(err error) *Exit {
	return &Exit{err: err}
}

func (e *Exit) Unwrap() error { return e.err }

func (e *Exit) Error() string {
	if e.err == nil {
		return e.msg
	}
	return fmt.Sprintf("%s: %s", e.msg, e.err.Error())
}
