package config

import (
	"encoding/json"
	"github.com/johnpoint/ControlCenter-Server/src/model"
	"io"
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	// 测试配置文件内容
	var testData = model.Config{AllowAddress: []string{"127.0.0.1"}, ListenPort: string("1323"), TLS: false, CERTPath: "PATHtoCER", KEYPath: "PATHtoKEY", Salt: "ControlCenter", Database: "testdata.db", RedisConfig: model.RedisConfig{Addr: "127.0.0.1:6379", Password: "", DB: 1}}
	file, _ := os.Create("config.json")
	defer file.Close()
	databy, _ := json.Marshal(testData)
	io.WriteString(file, string(databy)) // 写入测试配置文件
	var getConf model.Config
	getConf = LoadConfig()
	if getConf.AllowAddress[0] != testData.AllowAddress[0] {
		t.Error("AllowAddress Not Match")
	}
	if getConf.ListenPort != testData.ListenPort {
		t.Error("ListenPort Not Match")
	}
	if getConf.TLS != testData.TLS {
		t.Error("TLS Not Match")
	}
	if getConf.CERTPath != testData.CERTPath {
		t.Error("CERTPath Not Match")
	}
	if getConf.KEYPath != testData.KEYPath {
		t.Error("KEYPath Not Match")
	}
	if getConf.Salt != testData.Salt {
		t.Error("Salt Not Match")
	}
	if getConf.Database != testData.Database {
		t.Error("Database Not Match")
	}
}
