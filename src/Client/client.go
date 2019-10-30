package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

type Data struct {
	Base         DataBase
	Sites        []DataSite
	Certificates []DataCertificate
	Services     []DataService
}

type DataBase struct {
	ServerIpv4  string
	ServerIpv6  string
	HostName    string
	Token       string
	PollAddress string
}

type DataSite struct {
	Domain string
	Enable bool
	CerID  int64
}

type DataService struct {
	Name    string
	Enable  string
	Disable string
	Status  string
}

type DataCertificate struct {
	ID        int64
	Domain    string
	FullChain string
	Key       string
}

type StatusServer struct {
	Percent  StatusPercent
	CPU      []CPUInfo
	Mem      MemInfo
	Swap     SwapInfo
	Load     *load.AvgStat
	Network  map[string]InterfaceInfo
	BootTime uint64
	Uptime   uint64
}
type StatusPercent struct {
	CPU  float64
	Disk float64
	Mem  float64
	Swap float64
}
type CPUInfo struct {
	ModelName string
	Cores     int32
}
type MemInfo struct {
	Total     uint64
	Used      uint64
	Available uint64
}
type SwapInfo struct {
	Total     uint64
	Used      uint64
	Available uint64
}
type InterfaceInfo struct {
	Addrs    []string
	ByteSent uint64
	ByteRecv uint64
}

type Webreq struct {
	Code int64  `json:Code`
	Info string `json:Info`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("请输入参数")
		return
	}
	if os.Args[1] == "install" {
		if len(os.Args) != 5 {
			fmt.Println("参数数量错误错误")
			return
		}
		setup(os.Args)
		statuspoll()
		return
	}
	if os.Args[1] == "poll" {
		statuspoll()
		return
	}
	if os.Args[1] == "debug" {
		delSite("lvcshu.com")
		return
	}
	fmt.Println("未知的参数")
}

func setup(args []string) {
	data := Data{}
	url := args[2] + "/server/setup"
	method := "POST"

	payload := strings.NewReader("hostname=" + args[3] + "&ipv4=" + args[4])

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

func addCer(id int64, domain string, fullchain string, key string) bool {
	data := getData()
	for index := 0; index < len(data.Certificates); index++ {
		if data.Certificates[index].ID == id {
			fmt.Println("Certificate already exists")
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
	fmt.Println("OK!")
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
			fmt.Println("OK!")
			return true
		}
	}
	fmt.Println("Certificate not exists")
	return false
}

func statuspoll() {
	data := getData()
	fmt.Println("[ Poll start ] To " + data.Base.PollAddress)
	defer func() {
		fmt.Println("状态推送失败! 请检查服务端状态")
	}()
	for true {
		url := data.Base.PollAddress + "/server/update/" + data.Base.Token
		method := "POST"

		payload := strings.NewReader("ipv4=" + data.Base.ServerIpv4 + "&token=" + data.Base.Token + "&status=" + infoMiniJSON())

		client := &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		}
		req, err := http.NewRequest(method, url, payload)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		res, err := client.Do(req)
		if res != nil {
			fmt.Println("⇨ Poll Update To " + data.Base.PollAddress)
		}
		if err != nil {
			fmt.Println("状态推送失败! 请检查服务端状态")
			fmt.Println(err)
		}
		time.Sleep(time.Duration(2) * time.Second)
	}
}

func infoMiniJSON() string {
	v, _ := mem.VirtualMemory()
	s, _ := mem.SwapMemory()
	c, _ := cpu.Info()
	cc, _ := cpu.Percent(time.Second, false)
	d, _ := disk.Usage("/")
	n, _ := host.Info()
	nv, _ := net.IOCounters(true)
	l, _ := load.Avg()
	i, _ := net.Interfaces()
	ss := new(StatusServer)
	ss.Load = l
	ss.Uptime = n.Uptime
	ss.BootTime = n.BootTime
	ss.Percent.Mem = v.UsedPercent
	ss.Percent.CPU = cc[0]
	ss.Percent.Swap = s.UsedPercent
	ss.Percent.Disk = d.UsedPercent
	ss.CPU = make([]CPUInfo, len(c))
	for i, ci := range c {
		ss.CPU[i].ModelName = ci.ModelName
		ss.CPU[i].Cores = ci.Cores
	}
	ss.Mem.Total = v.Total
	ss.Mem.Available = v.Available
	ss.Mem.Used = v.Used
	ss.Swap.Total = s.Total
	ss.Swap.Available = s.Free
	ss.Swap.Used = s.Used
	ss.Network = make(map[string]InterfaceInfo)
	for _, v := range nv {
		var ii InterfaceInfo
		ii.ByteSent = v.BytesSent
		ii.ByteRecv = v.BytesRecv
		ss.Network[v.Name] = ii
	}
	for _, v := range i {
		if ii, ok := ss.Network[v.Name]; ok {
			ii.Addrs = make([]string, len(v.Addrs))
			for i, vv := range v.Addrs {
				ii.Addrs[i] = vv.Addr
			}
			ss.Network[v.Name] = ii
		}
	}
	b, err := json.Marshal(ss)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
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
