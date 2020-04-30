package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

func addCer(id int64, domain string, fullchain string, key string) bool {
	data := getData()
	for index := 0; index < len(data.Certificates); index++ {
		if data.Certificates[index].ID == id {
			log.Print("Certificate already exists")
			return false
		}
	}
	data.Certificates = append(data.Certificates, DataCertificate{ID: id, Domain: domain, FullChain: fullchain, Key: key})
	file, _ := os.Create("data.json")
	defer file.Close()
	databy, _ := json.Marshal(data)
	_, err := io.WriteString(file, string(databy))
	if err != nil {
		panic(err)
	}
	log.Print("OK!")
	return true
}

func syncCer() bool {
	sslPath := "/web/ssl/"
	data := getData()
	for i := 0; i < len(data.Certificates); i++ {
		if _, err := os.Stat(sslPath + data.Certificates[i].Domain); os.IsNotExist(err) {
			os.Mkdir(sslPath+data.Certificates[i].Domain, 0777)
		}
		fc, _ := os.Create(sslPath + data.Certificates[i].Domain + "/" + data.Certificates[i].Domain + ".fc") // .fc as fullchain
		defer fc.Close()
		_, err := io.WriteString(fc, data.Certificates[i].FullChain)
		if err != nil {
			panic(err)
		}
		key, _ := os.Create(sslPath + data.Certificates[i].Domain + "/" + data.Certificates[i].Domain + ".key")
		_, err = io.WriteString(key, data.Certificates[i].Key)
		if err != nil {
			panic(err)
		}
	}
	return true
}

func delCer(id int64) bool {
	data := getData()
	for index := 0; index < len(data.Certificates); index++ {
		if data.Certificates[index].ID == id {
			data.Certificates = append(data.Certificates[:index], data.Certificates[index+1:]...)
			file, _ := os.Create("data.json")
			defer file.Close()
			databy, _ := json.Marshal(data)
			_, err := io.WriteString(file, string(databy))
			if err != nil {
				panic(err)
			}
			log.Print("OK!")
			return true
		}
	}
	log.Print("Certificate not exists")
	return false
}
