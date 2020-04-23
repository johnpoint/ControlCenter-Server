package main

import (
	"encoding/json"
	"fmt"
	"github.com/docker/distribution/context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/inconshreveable/go-update"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

func poll() {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	var ver int64 = 0
	wg.Add(1)
	go listenUpdate(&mutex, &ver)
	go status(&mutex, &ver)
	wg.Wait()
}

func status(mutex *sync.Mutex, ver *int64) {
	data := getData()
	fmt.Println("[ Poll start ] To " + data.Base.PollAddress)
	defer func() {
		fmt.Println("状态推送失败! 请检查服务端状态")
	}()
	url := data.Base.PollAddress + "/server/update/" + data.Base.Token
	method := "POST"

	for true {
		if *ver == 1 {
			os.Exit(0)
		}
		mutex.Lock()
		payload := strings.NewReader("ipv4=" + data.Base.ServerIpv4 + "&token=" + data.Base.Token + "&status=" + infoMiniJSON())

		client := &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		}
		req, _ := http.NewRequest(method, url, payload)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		res, err := client.Do(req)
		if res != nil {
			defer res.Body.Close()
			fmt.Println(":: Poll Update To " + data.Base.PollAddress)
		}

		if err != nil {
			fmt.Println("状态推送失败! 请检查服务端状态")
			fmt.Println(err)
		}
		mutex.Unlock()
		time.Sleep(time.Duration(10) * time.Second)
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
	cli, err := client.NewEnvClient()
	defer cli.Close()
	if err != nil {
		fmt.Println(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All: true})
	if err != nil {
		fmt.Println(err)
	}

	for _, container := range containers {
		var str string
		for _, port := range container.Ports {
			str += strconv.FormatInt(int64(port.PrivatePort), 10) + " --> " + strconv.FormatInt(int64(port.PublicPort), 10) + "<br>"
		}

		docker := DockerInfo{}
		docker.Port = str
		docker.ID = container.ID
		docker.Name = container.Names[0]
		docker.Image = container.Image
		docker.State = container.Status
		ss.DockerInfo = append(ss.DockerInfo, docker)
	}
	b, err := json.Marshal(ss)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func listenUpdate(mutex *sync.Mutex, ver *int64) {
	data := getData()
	url := data.Base.PollAddress + "/server/now/" + data.Base.Token
	method := "GET"

	payload := strings.NewReader("ipv4=" + data.Base.ServerIpv4 + "&token=" + data.Base.Token + "&status=" + infoMiniJSON())

	for true {
		time.Sleep(time.Duration(2) * time.Second)
		mutex.Lock()
		client := &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		}
		req, _ := http.NewRequest(method, url, payload)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		res, err := client.Do(req)
		fmt.Println("Get Now Message")
		if res != nil {
			decoder := json.NewDecoder(res.Body)
			defer res.Body.Close()
			data := Webreq{}
			err := decoder.Decode(&data)
			if err != nil {
				fmt.Println("Error:", err)
				mutex.Unlock()
				continue
			}
			if data.Code == 211 {
				res.Body.Close()
				fmt.Println("Update to new version")
				resp, err := http.Get("https://cdn.lvcshu.info/xva/new/Client")
				if err != nil {
					fmt.Println(err)
					mutex.Unlock()
					continue
				}
				defer resp.Body.Close()
				err = update.Apply(resp.Body, update.Options{})
				if err != nil {
					fmt.Println(err)
					if rerr := update.RollbackError(err); rerr != nil {
						fmt.Println("Failed to rollback from bad update: %v", rerr)
					}
				}
				os.Chmod(os.Args[0], 0777)
				*ver = 1
				if err = syscall.Exec(os.Args[0], os.Args, os.Environ()); err != nil {
					panic(err)
				}
				mutex.Unlock()
				return
			}
			if data.Code == 210 {
				fmt.Println("Exit")
				os.Exit(0)
			}
			if data.Code == 212 {
				getUpdate()
				syncCer()
			}
		}

		if err != nil {
			fmt.Println("与服务端通信失败! 请检查服务端状态")
			fmt.Println(err)
		}
		mutex.Unlock()
	}
}
