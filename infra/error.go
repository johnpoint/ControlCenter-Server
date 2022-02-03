package infra

import "ControlCenter/pkg/errorhelper"

var (
	// 请求异常 401xx
	ReqParseError        = &errorhelper.Err{Code: 40100, Message: "请求参数异常"}
	ReqSameUsernameError = &errorhelper.Err{Code: 40101, Message: "有相同用户名"}

	// 认证异常 403xx
	ErrNeedLogin            = &errorhelper.Err{Code: 40301, Message: "此区域需要登录访问"}
	ErrNeedVerifyInfo       = &errorhelper.Err{Code: 40302, Message: "请求需要身份认证信息"}
	ErrAuthInfoInvalid      = &errorhelper.Err{Code: 40303, Message: "身份认证失败"}
	ErrAuthService          = &errorhelper.Err{Code: 40304, Message: "认证服务异常"}
	ErrAssetsAuthorityError = &errorhelper.Err{Code: 40305, Message: "资产获取异常"}
	ErrNoPermission         = &errorhelper.Err{Code: 40306, Message: "无权限进行操作"}

	// 数据库异常 501xx
	DataBaseError = &errorhelper.Err{Code: 50100, Message: "数据库异常"}
)
