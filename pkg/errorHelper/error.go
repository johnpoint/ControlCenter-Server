package errorHelper

import "errors"

type Code struct {
	Code    int
	Message string
	Error   error
}

var errMsgMap = map[int]string{
	-1: "Unknown error",
}

var errCodeMap = map[error]int{
	errors.New("nil"): -1,
}

func New(code int, msg string, err error) *Code {
	return &Code{
		Code:    code,
		Message: msg,
		Error:   err,
	}
}

func (e *Code) AddToMap() {
	if _, has := errMsgMap[e.Code]; has {
		panic("same err code")
	}
	errMsgMap[e.Code] = e.Message
	errCodeMap[e.Error] = e.Code
}

func GetErrMsg(err error) (int, string) {
	errCode, has := errCodeMap[err]
	if !has {
		errCode = -1
	}
	return errCode, errMsgMap[errCode]
}
