package xerror

import (
	"fmt"
	"runtime/debug"
)

type SeefsError struct {
	Inner      error
	Message    string
	StackTrace string
	Misc       map[string]any
}

// WrapError error with message
func WrapError(err error, messagef string, msgArgs ...any) SeefsError {
	return SeefsError{
		Inner:      err,
		Message:    fmt.Sprintf(messagef, msgArgs...),
		StackTrace: string(debug.Stack()),
		// stack message
		Misc: make(map[string]any),
	}
}

func (err SeefsError) Error() string {
	return err.Message
}
