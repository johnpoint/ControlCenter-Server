package apimiddleware

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var router *gin.Engine

func Get(uri string, router *gin.Engine) *httptest.ResponseRecorder {
	// 构造get请求
	req := httptest.NewRequest("GET", uri, nil)
	// 初始化响应
	w := httptest.NewRecorder()

	// 调用相应的handler接口
	router.ServeHTTP(w, req)
	return w
}

func TestLogPlusMiddleware(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	router = gin.New()
	router.Use(LogPlusMiddleware())
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"code": 200})
		return
	})
	var w *httptest.ResponseRecorder
	assert := assert.New(t)

	urlIndex := "/"
	w = Get(urlIndex, router)
	assert.Equal(200, w.Code)
}
