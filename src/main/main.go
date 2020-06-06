package main

import (
	"fmt"
	. "github.com/johnpoint/ControlCenter-Server/src/config"
	"github.com/johnpoint/ControlCenter-Server/src/router"
	"os"
)

func main() {
	if len(os.Args) == 2 {
		if os.Args[1] == "init" {
			InitServer()
		} else if os.Args[1] == "start" {
			router.Run()
		}
	}
	fmt.Println("参数错误")
}
