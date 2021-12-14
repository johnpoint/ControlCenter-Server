package infra

import "ControlCenter/pkg/errorHelper"

var (
	// 认证异常 403xxx
	ErrNeedLogin       = &errorHelper.Err{Code: 403001, Message: "此区域需要登录访问"}
	ErrNeedVerifyInfo  = &errorHelper.Err{Code: 403002, Message: "请求需要身份认证信息"}
	ErrAuthInfoInvalid = &errorHelper.Err{Code: 403003, Message: "身份认证失败"}
)
