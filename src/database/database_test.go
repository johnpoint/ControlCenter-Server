package database

import (
	"encoding/json"
	"fmt"
	. "github.com/johnpoint/ControlCenter-Server/src/model"
	"io"
	"os"
	"testing"
)

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
