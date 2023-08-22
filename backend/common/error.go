package common

import "fmt"

type CustomError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("message: %s, code: %d", e.Message, e.Code)
}

type CustomErrors struct {
	// 参数为空
	ParamEmptyError CustomError
	// 参数不合法
	ParamIllegalError CustomError
	// 意料之外
	UnexpectedError CustomError
}

var Errors = CustomErrors{
	ParamEmptyError:   CustomError{10105, "缺少关键参数"},
	ParamIllegalError: CustomError{10106, "参数类型错误"},
	UnexpectedError:   CustomError{99999, "发生意料之外的错误"},
}
