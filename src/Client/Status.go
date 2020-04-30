package main

import (
	"encoding/json"
	"github.com/docker/distribution/context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
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
	var timer int64 = 0
	data := getData()
	log.Print("[ Poll start ] To " + data.Base.PollAddress)
	for true {
		if timer == 10 {
			timer = 0
		}
		timer++
		time.Sleep(time.Duration(1) * time.Second)
		if timer%3 == 0 {
			defer func() {
				log.Print("状态推送失败! 请检查服务端状态")
			}()
			url := data.Base.PollAddress + "/server/update/" + data.Base.Token
			method := "POST"
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
				log.Print(":: Poll Update To " + data.Base.PollAddress)
			}

			if err != nil {
				log.Print("状态推送失败! 请检查服务端状态")
				log.Print(err)
			}
		}
		if timer%2 == 0 {
			data := getData()
			url := data.Base.PollAddress + "/server/now/" + data.Base.Token
			method := "GET"
			payload := strings.NewReader("ipv4=" + data.Base.ServerIpv4 + "&token=" + data.Base.Token + "&status=" + infoMiniJSON())
			client := &http.Client{
				CheckRedirect: func(req *http.Request, via []*http.Request) error {
					return http.ErrUseLastResponse
				},
			}
			req, _ := http.NewRequest(method, url, payload)
			req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
			res, err := client.Do(req)
			log.Print("Get Now Message")
			if res != nil {
				decoder := json.NewDecoder(res.Body)
				defer res.Body.Close()
				data := Webreq{}
				err := decoder.Decode(&data)
				if err != nil {
					log.Print("Error:", err)
					continue
				}
				if data.Code == 211 {
					res.Body.Close()
					log.Print("Update to new version")
					resp, err := http.Get("https://cdn.lvcshu.info/xva/new/Client")
					if err != nil {
						log.Print(err)
						continue
					}
					defer resp.Body.Close()
					err = update.Apply(resp.Body, update.Options{})
					if err != nil {
						log.Print(err)
						if rerr := update.RollbackError(err); rerr != nil {
							log.Print("Failed to rollback from bad update: %v", rerr)
						}
					}
					os.Chmod(os.Args[0], 0777)
					if err = syscall.Exec(os.Args[0], os.Args, os.Environ()); err != nil {
						panic(err)
					}
					return
				}
				if data.Code == 210 {
					log.Print("Exit")
					os.Exit(0)
				}
				if data.Code == 212 {
					getUpdate()
					syncCer()
				}
			}

			if err != nil {
				log.Print("与服务端通信失败! 请检查服务端状态")
				log.Print(err)
			}
		}
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
		log.Print(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All: true})
	if err != nil {
		log.Print(err)
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
	ss.Version = ClientVersion
	b, err := json.Marshal(ss)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}
