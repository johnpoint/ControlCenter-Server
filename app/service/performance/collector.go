package performance

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"time"
)

type Collector struct {
	Err              error                  `json:"-"`
	VirtualMemory    *mem.VirtualMemoryStat `json:"virtual_memory"`     // 内存
	SwapMemoryStat   *mem.SwapMemoryStat    `json:"swap_memory_stat"`   // 虚拟内存
	CPUStat          *CPUStat               `json:"cpu_stat"`           // CPU 信息
	DiskUsageStat    *disk.UsageStat        `json:"disk_usage_stat"`    // 硬盘信息
	SystemInfoStat   *host.InfoStat         `json:"system_info_stat"`   // 系统基本信息
	NetInterfaceStat []net.IOCountersStat   `json:"net_interface_stat"` // 网卡信息
	Load             *load.AvgStat          `json:"load"`               // 负载
	InterfaceStat    []net.InterfaceStat    `json:"interface_stat"`     // 网络接口信息
}

type CPUStat struct {
	Info    []cpu.InfoStat `json:"info"`
	Percent []float64      `json:"percent"`
}

func NewCollector() *Collector {
	return &Collector{}
}

func (c *Collector) getMem() {
	c.VirtualMemory, c.Err = mem.VirtualMemory()
	if c.Err != nil {
		return
	}
	c.SwapMemoryStat, c.Err = mem.SwapMemory()
	if c.Err != nil {
		return
	}
}

func (c *Collector) getCPU() {
	c.CPUStat.Info, c.Err = cpu.Info()
	if c.Err != nil {
		return
	}
	c.CPUStat.Percent, c.Err = cpu.Percent(time.Second, false)
	if c.Err != nil {
		return
	}
}

func (c *Collector) getDisk() {
	c.DiskUsageStat, c.Err = disk.Usage("/")
	if c.Err != nil {
		return
	}
}

func (c *Collector) getSysInfo() {
	c.SystemInfoStat, c.Err = host.Info()
	if c.Err != nil {
		return
	}
	c.Load, c.Err = load.Avg()
	if c.Err != nil {
		return
	}
}

func (c *Collector) GetNetInfo() {
	c.NetInterfaceStat, c.Err = net.IOCounters(true)
	if c.Err != nil {
		return
	}
	c.InterfaceStat, c.Err = net.Interfaces()
	if c.Err != nil {
		return
	}
}
