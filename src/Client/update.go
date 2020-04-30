package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

func getUpdate() bool {
	data := getData()
	url := data.Base.PollAddress + "/server/update/" + data.Base.Token
	method := "GET"
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	req, err := http.NewRequest(method, url, nil)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, err := client.Do(req)
	if err != nil {
		log.Print("状态获取失败! 请检查服务端状态")
		log.Print(err)
		return false
	}
	if res != nil {
		log.Print(":: Get update from " + data.Base.PollAddress)
		decoder := json.NewDecoder(res.Body)
		Getdata := UpdateInfo{}
		err := decoder.Decode(&Getdata)
		if err != nil {
			log.Print("Error:", err)
		}
		data.Certificates = Getdata.Certificates
		data.Sites = Getdata.Sites
		data.Docker = Getdata.Docker
		file, _ := os.Create("data.json")
		defer file.Close()
		databy, _ := json.Marshal(data)
		io.WriteString(file, string(databy))
		if err != nil {
			panic(err)
		}
		log.Print("OK!")
		return true
	}
	return false
}
