package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func addService(name string, enable string, disable string) bool {
	data := getData()
	for index := 0; index < len(data.Services); index++ {
		if data.Services[index].Name == name {
			fmt.Println("Service already exists")
			return false
		}
	}
	data.Services = append(data.Services, DataService{Name: name, Enable: enable, Disable: disable, Status: "stop"})
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

func delService(name string) bool {
	data := getData()
	for index := 0; index < len(data.Services); index++ {
		if data.Services[index].Name == name {
			data.Services = append(data.Services[:index], data.Services[index+1:]...)
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
	fmt.Println("Service not exists")
	return false
}
