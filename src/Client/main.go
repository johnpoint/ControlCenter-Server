package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
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
		statuspoll()
		return
	}
	if os.Args[1] == "help" {
		showhelp()
		return
	}
	if os.Args[1] == "poll" {
		statuspoll()
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
	fmt.Println("XvA Server Slave Client")
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
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	webreq := Webreq{}
	json.Unmarshal([]byte(body), &webreq)
	if webreq.Code != 200 {
		fmt.Println(webreq.Info)
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
	fmt.Println("OK!")
}

func addSite(domain string, enable bool, cerid int64) bool {
	data := getData()
	for index := 0; index < len(data.Sites); index++ {
		if data.Sites[index].Domain == domain {
			fmt.Println("Site already exists")
			return false
		}
	}
	data.Sites = append(data.Sites, DataSite{Domain: domain, Enable: enable, CerID: cerid})
	file, _ := os.Create("data.json")
	defer file.Close()
	databy, _ := json.Marshal(data)
	_, err := io.WriteString(file, string(databy))
	if err != nil {
		panic(err)
	}
	fmt.Println("OK!")
	return true
}

func delSite(domain string) bool {
	data := getData()
	for index := 0; index < len(data.Sites); index++ {
		if data.Sites[index].Domain == domain {
			data.Sites = append(data.Sites[:index], data.Sites[index+1:]...)
			file, _ := os.Create("data.json")
			defer file.Close()
			databy, _ := json.Marshal(data)
			_, err := io.WriteString(file, string(databy))
			if err != nil {
				panic(err)
			}
			fmt.Println("OK!")
			return true
		}
	}
	fmt.Println("Site not exists")
	return false
}

func getData() Data {
	file, _ := os.Open("data.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	data := Data{}
	err := decoder.Decode(&data)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return data
}
