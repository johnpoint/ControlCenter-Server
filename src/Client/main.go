package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

const ClientVersion = "1.9.4"

func main() {
	go func() {
		http.ListenAndServe("0.0.0.0:8080", nil)
	}()
	fmt.Println("ControlCenter-Client v", ClientVersion)
	if len(os.Args) < 2 {
		fmt.Println("请输入参数(help 可以调用帮助信息)")
		return
	}
	if os.Args[1] == "install" {
		if len(os.Args) != 7 {
			fmt.Println("参数数量错误")
			return
		}
		setup(os.Args)
		poll()
		return
	}
	if os.Args[1] == "help" {
		showhelp()
		return
	}
	if os.Args[1] == "poll" {
		poll()
		return
	}
	if os.Args[1] == "update" {
		getUpdate()
		return
	}
	if os.Args[1] == "sync" {
		syncCer()
		return
	}
	fmt.Println("未知的参数")
}

func showhelp() {
	fmt.Println("ControlCenter - Server Slave Client")
	fmt.Println("参数:")
	fmt.Println("poll - 开始向控制中心服务器推送状态")
	fmt.Println("update - 向控制中心服务器获取控制信息")
	fmt.Println("sync - 应用控制信息")
	fmt.Println("")
	fmt.Println("注册:")
	fmt.Println("install 控制中心服务器地址 `hostname` `curl ip.sb -4` `curl ip.sb` user_token")
}

func setup(args []string) {
	data := Data{}
	url := args[2] + "/server/setup/" + args[6]
	method := "POST"

	payload := strings.NewReader("hostname=" + args[3] + "&ipv4=" + args[4] + "&ipv6=" + args[5])

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		log.Print(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	webreq := Webreq{}
	err = json.Unmarshal([]byte(body), &webreq)
	if err != nil {
		log.Print(err)
	}
	if webreq.Code != 200 {
		log.Print(webreq.Info)
		return
	}
	base := DataBase{ServerIpv4: args[4], HostName: args[3], Token: webreq.Info, PollAddress: args[2]}
	data.Base = base
	file, _ := os.Create("data.json")
	defer file.Close()
	databy, err := json.Marshal(data)
	_, err1 := io.WriteString(file, string(databy))
	if err1 != nil {
		panic(err1)
	}
	log.Print("OK!")
}

func getData() Data {
	file, _ := os.Open("data.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	data := Data{}
	err := decoder.Decode(&data)
	if err != nil {
		log.Print("Error:", err)
	}
	return data
}
