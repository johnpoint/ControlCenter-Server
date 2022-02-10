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
	Data *Data
	Err  error `json:"-"`
}

type Data struct {
	VirtualMemory    *mem.VirtualMemoryStat `json:"virtual_memory"`     // 内存
	SwapMemoryStat   *mem.SwapMemoryStat    `json:"swap_memory_stat"`   // 虚拟内存
	CPUStat          *CPUStat               `json:"cpu_stat"`           // CPU 信息
	DiskUsageStat    *disk.UsageStat        `json:"disk_usage_stat"`    // 硬盘信息
	PartitionStat    []disk.PartitionStat   `json:"partition_stat"`     // 分区信息
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

func (c *Collector) Do() (*Data, error) {
	c.Data = &Data{
		CPUStat: &CPUStat{},
	}
	var steps = []func() error{
		c.getMem,
		c.getCPU,
		c.getDisk,
		c.getSysInfo,
		c.getNetInfo,
	}

	for i := range steps {
		if steps[i]() != nil {
			return nil, c.Err
		}
	}
	return c.Data, nil
}

func (c *Collector) getMem() error {
	c.Data.VirtualMemory, c.Err = mem.VirtualMemory()
	if c.Err != nil {
		return c.Err
	}
	c.Data.SwapMemoryStat, c.Err = mem.SwapMemory()
	if c.Err != nil {
		return c.Err
	}
	return nil
}

func (c *Collector) getCPU() error {
	c.Data.CPUStat.Info, c.Err = cpu.Info()
	if c.Err != nil {
		return c.Err
	}
	c.Data.CPUStat.Percent, c.Err = cpu.Percent(time.Second, false)
	if c.Err != nil {
		return c.Err
	}
	return nil
}

func (c *Collector) getDisk() error {
	c.Data.DiskUsageStat, c.Err = disk.Usage("/")
	if c.Err != nil {
		return c.Err
	}
	c.Data.PartitionStat, c.Err = disk.Partitions(true)
	if c.Err != nil {
		return c.Err
	}
	if len(c.Data.PartitionStat) > 30 {
		c.Data.PartitionStat = make([]disk.PartitionStat, 0)
	}
	return nil
}

func (c *Collector) getSysInfo() error {
	c.Data.SystemInfoStat, c.Err = host.Info()
	if c.Err != nil {
		return c.Err
	}
	c.Data.Load, c.Err = load.Avg()
	if c.Err != nil {
		return c.Err
	}
	return nil
}

func (c *Collector) getNetInfo() error {
	c.Data.NetInterfaceStat, c.Err = net.IOCounters(true)
	if c.Err != nil {
		return c.Err
	}
	c.Data.InterfaceStat, c.Err = net.Interfaces()
	if c.Err != nil {
		return c.Err
	}
	return nil
}
