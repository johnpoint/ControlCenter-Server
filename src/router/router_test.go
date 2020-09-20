package router

import (
	"ControlCenter-Server/src/apis"
	"ControlCenter-Server/src/database"
	"ControlCenter-Server/src/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	fmt.Println("==== 初始化资源 ===")
	del := os.Remove("./testdata.db")
	if del != nil {
		fmt.Println("= 数据库已经初始化 =")
	}
	var testData = model.Config{AllowAddress: []string{"127.0.0.1"}, ListenPort: string("1323"), TLS: false, CERTPath: "PATHtoCER", KEYPath: "PATHtoKEY", Salt: "ControlCenter", Database: "testdata.db", RedisConfig: struct {
		Enable   bool
		Addr     string
		Password string
		DB       int
	}{Addr: "127.0.0.1:6379", Password: "", DB: 1, Enable: true}}
	file, _ := os.Create("config.json")
	fmt.Println("= 配置文件设置完成 =")
	defer file.Close()
	databy, _ := json.Marshal(testData)
	io.WriteString(file, string(databy)) // 写入测试配置文件
	result := m.Run()                    //运行go的测试
	fmt.Println("=== 释放资源 ===")
	del = os.Remove("./testdata.db")
	if del != nil {
		fmt.Println(del)
	}
	os.Exit(result) //退出程序
}

func TestOnline(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if assert.NoError(t, apis.Accessible(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestLogin(t *testing.T) {
	e := echo.New()
	if !database.AddUser(model.User{ID: 233, Username: "testUser", Mail: "i@test.com", Password: "23333", Level: 1, Token: "23456"}) {
		panic("add user fail")
	}
	f := make(url.Values)
	f.Set("password", "23333")
	f.Set("email", "i@test.com")
	req := httptest.NewRequest(echo.POST, "/user/auth/login", strings.NewReader(f.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if assert.NoError(t, apis.OaLogin(c)) {
		fmt.Println(rec.Body)
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
