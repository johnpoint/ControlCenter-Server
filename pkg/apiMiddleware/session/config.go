package session

type SessionConfig struct {
	ExpireTime int         `json:"expire_time"` // 过期时间
	AutoRenew  bool        `json:"auto_renew"`  // 是否自动续期
	Prefix     string      `json:"prefix"`      // key前缀
	HeaderName string      `json:"header_name"` // session 请求头字段名
	CtxKey     string      `json:"ctx_key"`     // 提取出来的数据放在上下文的 key
	ReturnData interface{} `json:"-"`
}
