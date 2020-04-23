package main

import "github.com/shirou/gopsutil/load"

type (
	// Data model
	Data struct {
		Base         DataBase
		Sites        []DataSite
		Certificates []DataCertificate
		Docker       []DockerConfig
	}

	// DataBase model
	DataBase struct {
		ServerIpv4  string
		ServerIpv6  string
		HostName    string
		Token       string
		PollAddress string
	}

	// UpdateInfo model
	UpdateInfo struct {
		Code         int64
		Sites        []DataSite
		Certificates []DataCertificate
		Docker       []DockerConfig
	}

	DockerConfig struct {
		ID     int64
		Config string
		Status int64 //-1 关闭 | 1 开启
	}

	// DataSite model
	DataSite struct {
		ID     int64
		Domain string
		CerID  int64
		Config string
	}

	// DataDocker model
	DockerInfo struct {
		ID     string
		Name   string
		Status int64
		Port   string
		State  string
		Image  string
	}

	// DataCertificate model
	DataCertificate struct {
		ID        int64
		Domain    string
		FullChain string
		Key       string
	}

	// StatusServer model
	StatusServer struct {
		Version    string
		Percent    StatusPercent
		CPU        []CPUInfo
		Mem        MemInfo
		Swap       SwapInfo
		Load       *load.AvgStat
		Network    map[string]InterfaceInfo
		BootTime   uint64
		Uptime     uint64
		DockerInfo []DockerInfo
	}

	// StatusPercent model
	StatusPercent struct {
		CPU  float64
		Disk float64
		Mem  float64
		Swap float64
	}

	// CPUInfo model
	CPUInfo struct {
		ModelName string
		Cores     int32
	}

	// MemInfo model
	MemInfo struct {
		Total     uint64
		Used      uint64
		Available uint64
	}

	// SwapInfo model
	SwapInfo struct {
		Total     uint64
		Used      uint64
		Available uint64
	}

	// InterfaceInfo model
	InterfaceInfo struct {
		Addrs    []string
		ByteSent uint64
		ByteRecv uint64
	}

	// Webreq model
	Webreq struct {
		Code int64  `json:Code`
		Info string `json:Info`
	}
)
