package main

import (
	"fmt"
	"github.com/johnpoint/ControlCenter-Server/src/config"
	"github.com/johnpoint/ControlCenter-Server/src/router"
	"os"
)

func main() {
	if len(os.Args) == 2 {
		switch os.Args[1] {
		case "start":
			router.Run()
			break
		case "init":
			config.InitServer()
			break
			// TODO: 交互初始化配置文件
		case "test":
			fmt.Println("build pass")
			break
		case "update":
			config.UpdateConfig()
			break
		default:
			fmt.Println("参数错误")
		}
	}
	return
}
