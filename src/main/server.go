package main

import (
	"fmt"
	"github.com/johnpoint/ControlCenter-Server/src/config"
	"github.com/johnpoint/ControlCenter-Server/src/router"
	"os"
)

func main() {
	if len(os.Args) == 2 {
		if os.Args[1] == "init" {
			config.InitServer()
			// TODO: 交互初始化配置文件
		} else if os.Args[1] == "start" {
			router.Run()
		} else if os.Args[1] == "test" {
			fmt.Println("build pass")
			return
		}
	}
	fmt.Println("参数错误")
	return
}
