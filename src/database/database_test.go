package database

import (
	"encoding/json"
	"fmt"
	"github.com/johnpoint/ControlCenter-Server/src/model"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("==== 初始化资源 ===")
	del := os.Remove("./testdata.db")
	if del != nil {
		fmt.Println("数据库已经初始化")
	}
	result := m.Run() //运行go的测试
	fmt.Println("=== 释放资源 ===")
	del = os.Remove("./testdata.db")
	if del != nil {
		panic(del)
	}
	os.Exit(result) //退出程序
}

func TestInitDatabase(t *testing.T) {
	// 测试配置文件内容
	var testData = model.Config{AllowAddress: []string{"127.0.0.1"}, ListenPort: string("1323"), TLS: false, CERTPath: "PATHtoCER", KEYPath: "PATHtoKEY", Salt: "ControlCenter", Database: "testdata.db", RedisConfig: struct {
		Enable   bool
		Addr     string
		Password string
		DB       int
	}{Addr: "127.0.0.1:6379", Password: "", DB: 1, Enable: true}}
	testData.Debug = true
	file, _ := os.Create("config.json")
	defer file.Close()
	databy, _ := json.Marshal(testData)
	io.WriteString(file, string(databy)) // 写入测试配置文件
	if initDatabase() == nil {
		t.Fatal("connect db error")
	}
	testData.Debug = false
	file, _ = os.Create("config.json")
	defer file.Close()
	databy, _ = json.Marshal(testData)
	io.WriteString(file, string(databy)) // 写入测试配置文件
	if initDatabase() == nil {
		t.Fatal("connect db error")
	}
}

func TestAddServer(t *testing.T) {
	if !AddServer(model.Server{ID: 10086, Hostname: "testServer", Ipv4: "8.8.8.8", Ipv6: "::1", UID: 1, Token: "TestToken", Online: 1}) {
		t.Fatal("add server1 fail")
	}
	if !AddServer(model.Server{ID: 10087, Hostname: "testServer", Ipv4: "1.1.1.1", Ipv6: "::2", UID: 1, Token: "TestToken", Online: 1}) {
		t.Fatal("add server1 fail")
	}
	if AddServer(model.Server{ID: 10087, Hostname: "testServer", Ipv4: "1.1.1.1", Ipv6: "::2", UID: 1, Token: "TestToken", Online: 1}) {
		t.Fatal("add server1 fail")
	}
}

func TestGetServer(t *testing.T) {
	ServerInfo := GetServer(model.Server{Hostname: "testServer"})
	want := []model.Server{}
	want = append(want, model.Server{ID: 10086, Hostname: "testServer", Ipv4: "8.8.8.8", Ipv6: "::1", UID: 1, Token: "TestToken", Online: 1}, model.Server{ID: 10087, Hostname: "testServer", Ipv4: "1.1.1.1", Ipv6: "::2", UID: 1, Token: "TestToken", Online: 1})
	assert.Equal(t, want, ServerInfo)
}

func TestDelServer(t *testing.T) {
	if !DelServer(10086, 1) {
		panic("del server fail")
	}
	ServerInfo := GetServer(model.Server{Hostname: "testServer"})
	want := []model.Server{}
	want = append(want, model.Server{ID: 10087, Hostname: "testServer", Ipv4: "1.1.1.1", Ipv6: "::2", UID: 1, Token: "TestToken", Online: 1})
	assert.Equal(t, want, ServerInfo)
}

func TestUpdateServer(t *testing.T) {
	if !UpdateServer(model.Server{ID: 10087}, model.Server{Hostname: "Server"}) {
		panic("update server fail")
	}
	ServerInfo := GetServer(model.Server{ID: 10087})
	want := []model.Server{}
	want = append(want, model.Server{ID: 10087, Hostname: "Server", Ipv4: "1.1.1.1", Ipv6: "::2", UID: 1, Token: "TestToken", Online: 1})
	assert.Equal(t, want, ServerInfo)
}

func TestAddUser(t *testing.T) {
	if !AddUser(model.User{ID: 233, Username: "testUser", Mail: "i@test.com", Password: "23333", Level: 1, Token: "23456"}) {
		panic("add user fail")
	}
	user := GetUser(model.User{ID: 233})
	want := []model.User{}
	want = append(want, model.User{ID: 233, Username: "testUser", Mail: "i@test.com", Password: "23333", Level: 1, Token: "23456"})
	assert.Equal(t, want, user)
}

func TestUpdateUser(t *testing.T) {
	if !UpdateUser(model.User{ID: 233}, model.User{Username: "User"}) {
		panic("add user fail")
	}
	user := GetUser(model.User{ID: 233})
	want := []model.User{}
	want = append(want, model.User{ID: 233, Username: "User", Mail: "i@test.com", Password: "23333", Level: 1, Token: "23456"})
	assert.Equal(t, want, user)
}

func TestDelUser(t *testing.T) {
	if !DelUser(model.User{ID: 233}) {
		panic("del user fail")
	}
	user := GetUser(model.User{})
	want := []model.User{}
	assert.Equal(t, want, user)
}
