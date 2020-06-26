package database

import (
	"encoding/json"
	"fmt"
	. "github.com/johnpoint/ControlCenter-Server/src/model"
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
	var testData = Config{AllowAddress: []string{"127.0.0.1"}, ListenPort: string("1323"), TLS: false, CERTPath: "PATHtoCER", KEYPath: "PATHtoKEY", Salt: "ControlCenter", Database: "testdata.db"}
	file, _ := os.Create("config.json")
	defer file.Close()
	databy, _ := json.Marshal(testData)
	io.WriteString(file, string(databy)) // 写入测试配置文件
	if initDatabase() == nil {
		t.Fatal("connect db error")
	}
}

func TestAddServer(t *testing.T) {
	if !AddServer(Server{ID: 10086, Hostname: "testServer", Ipv4: "8.8.8.8", Ipv6: "::1", UID: 1, Token: "TestToken", Online: 1, Update: 1}) {
		t.Fatal("add server1 fail")
	}
	if !AddServer(Server{ID: 10087, Hostname: "testServer", Ipv4: "1.1.1.1", Ipv6: "::2", UID: 1, Token: "TestToken", Online: 1, Update: 1}) {
		t.Fatal("add server1 fail")
	}
}

func ExampleGetServer() {
	ServerInfo := GetServer(Server{Hostname: "testServer"})
	fmt.Println(ServerInfo)
	// Output: [{ testServer 8.8.8.8 ::1 10086 1 TestToken 1 1} { testServer 1.1.1.1 ::2 10087 1 TestToken 1 1}]
}

func ExampleDelServer() {
	if !DelServer(10086, 1) {
		panic("del server fail")
	}
	ServerInfo := GetServer(Server{Hostname: "testServer"})
	fmt.Println(ServerInfo)
	// Output: [{ testServer 1.1.1.1 ::2 10087 1 TestToken 1 1}]
}

func ExampleUpdateServer() {
	if !UpdateServer(Server{ID: 10087}, Server{Hostname: "Server"}) {
		panic("update server fail")
	}
	ServerInfo := GetServer(Server{ID: 10087})
	fmt.Println(ServerInfo)
	// Output: [{ Server 1.1.1.1 ::2 10087 1 TestToken 1 1}]
}

func ExampleAddUser() {
	if !AddUser(User{ID: 233, Username: "testUser", Mail: "i@test.com", Password: "23333", Level: 1, Token: "23456"}) {
		panic("add user fail")
	}
	user := GetUser(User{ID: 233})
	fmt.Println(user)
	// Output: [{233 testUser i@test.com 23333 1 23456}]
}

func ExampleUpdateUser() {
	if !UpdateUser(User{ID: 233}, User{Username: "User"}) {
		panic("add user fail")
	}
	user := GetUser(User{ID: 233})
	fmt.Println(user)
	// Output: [{233 User i@test.com 23333 1 23456}]
}

func ExampleDelUser() {
	if !DelUser(User{ID: 233}) {
		panic("del user fail")
	}
	user := GetUser(User{})
	fmt.Println(user)
	// Output: []
}
