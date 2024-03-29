package apiMiddleware

import (
	"ControlCenter/infra"
	"ControlCenter/pkg/errorhelper"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"strings"
	"time"
)

// JwtClaims 自定义 data 数据结构
type JwtClaims struct {
	UUID string `json:"uuid"`
	jwt.StandardClaims
}

var JwtTokenExpireDuration = time.Hour * 24 * 7 // jwt token 过期时间（默认7天）
var JwtSecret = []byte("JwtSecret")

// GenToken 生成JWT
func GenToken(uuid string) (string, error) {
	// 创建一个我们自己的声明
	c := JwtClaims{
		uuid, // 用户ID
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(JwtTokenExpireDuration).Unix(), // 过期时间
			Issuer:    "lvcshu",                                      // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(JwtSecret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*JwtClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return JwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": infra.ErrNeedVerifyInfo,
				"msg":  errorhelper.GetErrMessage(infra.ErrNeedVerifyInfo),
			})
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusOK, gin.H{
				"code": infra.ErrNeedVerifyInfo,
				"msg":  errorhelper.GetErrMessage(infra.ErrNeedVerifyInfo),
			})
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": infra.ErrAuthInfoInvalid,
				"msg":  errorhelper.GetErrMessage(infra.ErrAuthInfoInvalid),
			})
			c.Abort()
			return
		}
		if mc.ExpiresAt-time.Now().Unix() < 24*60*60 {
			token, _ := GenToken(mc.UUID)
			c.SetCookie("jwt", token, 0, "", "", true, true)
		}
		// 将当前请求的username信息保存到请求的上下文c上
		c.Set("uid", mc.UUID)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}
