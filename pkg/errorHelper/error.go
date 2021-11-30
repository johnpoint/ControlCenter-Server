package errorHelper

import (
	"fmt"
)

var (
	// 通用错误码
	OK      = &Err{Code: 0, Message: "OK"}
	Unknown = &Err{Code: -1, Message: "未知错误"}

	// 认证异常 403xxx
	ErrNeedLogin      = &Err{Code: 403001, Message: "此区域需要登录访问"}
	ErrNeedVerifyInfo = &Err{Code: 403002, Message: "请求需要身份认证信息"}
)

// Err 定义错误
type Err struct {
	Code      int    // 错误码
	Message   string // 展示给用户看的
	ErrorInfo error  // 保存内部错误信息
}

func (err *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, error: %s", err.Code, err.Message, err.ErrorInfo)
}

func GetErrCode(err error) int {
	trueErr, ok := err.(*Err)
	if !ok {
		return Unknown.Code
	}
	return trueErr.Code
}

func GetErrMessage(err error) string {
	trueErr, ok := err.(*Err)
	if !ok {
		return Unknown.Message
	}
	return trueErr.Message
}

func DecodeErr(err error) (int, string) {
	if err == nil {
		return OK.Code, OK.Message
	}

	trueErr, ok := err.(*Err)
	if !ok {
		return Unknown.Code, Unknown.Message
	}
	return trueErr.Code, trueErr.Message
}
