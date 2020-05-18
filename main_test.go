package main

import (
	"crypto/md5"
	"fmt"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	conf     = loadConfig()
	salt     = conf.Salt
	userJSON = `{"email":"i@test.com","password":"123456passwword"}`
)

func TestoaLogin(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/user/auth/login", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	u := User{ID: 233, Mail: "i@test.com", Password: "123456password"}
	data := []byte(u.Mail + salt + u.Password)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has)
	h := User{ID: 233, Mail: "i@test.com", Password: md5str1, Level: 0, Token: "123456"}
	addUser(h)
	if assert.NoError(t, nil) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, userJSON, rec.Body.String())
	}
}
