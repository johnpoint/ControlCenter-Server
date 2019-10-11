package main

import (
	"encoding/json"
	"fmt"
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

type Configuration struct {
	PollAddress  string
	ServerToken  string
	DataFilePath string
}

type Data struct {
	Base         DataBase
	Sites        []DataSite
	Certificates []DataCertificate
}

type DataBase struct {
	ServerIpv4 string
	ServerIpv6 string
	HostName   string
	Token      string
}

type DataSite struct {
	Domain string
	Enable bool
	CerID  int64
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

func main() {
	if len(os.Args) < 2 {
		fmt.Println("请输入参数")
		return
	}
	if os.Args[1] == "install" {
		//install `hostname` `curl ip.sb`
		setup(os.Args)
		return
	}
	if os.Args[1] == "poll" {
		statuspoll()
		return
	}
	if os.Args[1] == "data" {
		fmt.Println(getData())
		return
	}
	fmt.Println("未知的参数")
}

func setup(args []string) {
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

	fmt.Println(string(body))
}

func statuspoll() {
	data := getData()
	conf := loadConfig()
	fmt.Println("[ Poll start ] To " + conf.PollAddress)
	for true {
		url := conf.PollAddress + "/server/update/" + conf.ServerToken
		method := "POST"

		payload := strings.NewReader("ipv4=" + data.Base.ServerIpv4 + "&token=" + data.Base.Token + "&status=" + infoMiniJSON())

		client := &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		}
		req, err := http.NewRequest(method, url, payload)

		if err != nil {
			fmt.Println("状态推送失败! 请检查服务端状态")
			fmt.Println(err)
		}
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		res, err := client.Do(req)
		if err != nil {
			fmt.Println("状态推送失败! 请检查服务端状态")
			fmt.Println(err)
		}
		if res.Status != "200 OK" {
			fmt.Println("状态推送失败! 请检查配置")
			break
		}
		defer res.Body.Close()
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

func loadConfig() Configuration {
	file, _ := os.Open("config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	conf := Configuration{}
	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return conf
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
