package xerror

import (
	"fmt"
	"runtime/debug"
)

type SeefsError struct {
	Inner      error
	Message    string
	StackTrace string
	Misc       map[string]interface{}
}

// WrapError error包装
func WrapError(err error, messagef string, msgArgs ...interface{}) SeefsError {
	return SeefsError{
		Inner:      err,
		Message:    fmt.Sprintf(messagef, msgArgs...),
		StackTrace: string(debug.Stack()),
		// 堆栈跟踪的hash或可能有助于诊断错误的其他上下文信息
		Misc: make(map[string]interface{}),
	}
}

func (err SeefsError) Error() string {
	return err.Message
}
