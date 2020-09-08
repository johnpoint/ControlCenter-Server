package database

import (
	"encoding/json"
	"fmt"
	"github.com/johnpoint/ControlCenter-Server/src/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"io"
	"os"
	"reflect"
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

func TestAddCer(t *testing.T) {
	type args struct {
		certificate model.Certificate
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddCer(tt.args.certificate); got != tt.want {
				t.Errorf("AddCer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddEvent(t *testing.T) {
	type args struct {
		eventType int64
		target    int64
		code      int64
		info      string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddEvent(tt.args.eventType, tt.args.target, tt.args.code, tt.args.info); got != tt.want {
				t.Errorf("AddEvent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddLog(t *testing.T) {
	type args struct {
		service string
		event   string
		level   int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddLog(tt.args.service, tt.args.event, tt.args.level); got != tt.want {
				t.Errorf("AddLog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddServer1(t *testing.T) {
	type args struct {
		server model.Server
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddServer(tt.args.server); got != tt.want {
				t.Errorf("AddServer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddUser1(t *testing.T) {
	type args struct {
		user model.User
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddUser(tt.args.user); got != tt.want {
				t.Errorf("AddUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDelCer(t *testing.T) {
	type args struct {
		certificate model.Certificate
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DelCer(tt.args.certificate); got != tt.want {
				t.Errorf("DelCer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDelServer1(t *testing.T) {
	type args struct {
		id  int64
		uid int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DelServer(tt.args.id, tt.args.uid); got != tt.want {
				t.Errorf("DelServer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDelUser1(t *testing.T) {
	type args struct {
		user model.User
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DelUser(tt.args.user); got != tt.want {
				t.Errorf("DelUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFinishEvent(t *testing.T) {
	type args struct {
		id int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FinishEvent(tt.args.id); got != tt.want {
				t.Errorf("FinishEvent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCer(t *testing.T) {
	type args struct {
		certificate model.Certificate
	}
	tests := []struct {
		name string
		args args
		want []model.Certificate
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCer(tt.args.certificate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetConfig(t *testing.T) {
	type args struct {
		config model.SysConfig
	}
	tests := []struct {
		name string
		args args
		want []model.SysConfig
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetConfig(tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetEvent(t *testing.T) {
	type args struct {
		eventType int64
		target    int64
		code      int64
		info      string
		active    int64
	}
	tests := []struct {
		name string
		args args
		want []model.Event
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetEvent(tt.args.eventType, tt.args.target, tt.args.code, tt.args.info, tt.args.active); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEvent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetLinkCer(t *testing.T) {
	type args struct {
		serverLink model.ServerLink
	}
	tests := []struct {
		name string
		args args
		want []model.ServerLink
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetServerLinkedItem(tt.args.serverLink); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetServerLinkedItem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetServer1(t *testing.T) {
	type args struct {
		server model.Server
	}
	tests := []struct {
		name string
		args args
		want []model.Server
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetServer(tt.args.server); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetServer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetUser(t *testing.T) {
	type args struct {
		user model.User
	}
	tests := []struct {
		name string
		args args
		want []model.User
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetUser(tt.args.user); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkServer(t *testing.T) {
	type args struct {
		serverLink model.ServerLink
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LinkServer(tt.args.serverLink); got != tt.want {
				t.Errorf("LinkServer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetConfig(t *testing.T) {
	type args struct {
		config model.SysConfig
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetConfig(tt.args.config); got != tt.want {
				t.Errorf("SetConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnLinkServer(t *testing.T) {
	type args struct {
		serverLink model.ServerLink
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UnLinkServer(tt.args.serverLink); got != tt.want {
				t.Errorf("UnLinkServer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateCer(t *testing.T) {
	type args struct {
		where       model.Certificate
		certificate model.Certificate
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UpdateCer(tt.args.where, tt.args.certificate); got != tt.want {
				t.Errorf("UpdateCer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateServer1(t *testing.T) {
	type args struct {
		where  model.Server
		server model.Server
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UpdateServer(tt.args.where, tt.args.server); got != tt.want {
				t.Errorf("UpdateServer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateUser1(t *testing.T) {
	type args struct {
		where model.User
		user  model.User
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UpdateUser(tt.args.where, tt.args.user); got != tt.want {
				t.Errorf("UpdateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_addConfig(t *testing.T) {
	type args struct {
		config model.SysConfig
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addConfig(tt.args.config); got != tt.want {
				t.Errorf("addConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_initDatabase(t *testing.T) {
	tests := []struct {
		name string
		want *gorm.DB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := initDatabase(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("initDatabase() = %v, want %v", got, tt.want)
			}
		})
	}
}
