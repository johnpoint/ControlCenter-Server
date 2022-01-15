package infra

import "ControlCenter/pkg/errorHelper"

var (
	// 请求异常 401xx
	ReqParseError        = &errorHelper.Err{Code: 40100, Message: "请求参数异常"}
	ReqSameUsernameError = &errorHelper.Err{Code: 40101, Message: "有相同用户名"}

	// 认证异常 403xx
	ErrNeedLogin            = &errorHelper.Err{Code: 40301, Message: "此区域需要登录访问"}
	ErrNeedVerifyInfo       = &errorHelper.Err{Code: 40302, Message: "请求需要身份认证信息"}
	ErrAuthInfoInvalid      = &errorHelper.Err{Code: 40303, Message: "身份认证失败"}
	ErrAuthService          = &errorHelper.Err{Code: 40304, Message: "认证服务异常"}
	ErrAssetsAuthorityError = &errorHelper.Err{Code: 40305, Message: "资产获取异常"}

	// 数据库异常 501xx
	DataBaseError = &errorHelper.Err{Code: 50100, Message: "数据库异常"}
)
