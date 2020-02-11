package main

import "github.com/shirou/gopsutil/load"

// Data model
type Data struct {
	Base         DataBase
	Sites        []DataSite
	Certificates []DataCertificate
	Services     []DataService
}

// DataBase model
type DataBase struct {
	ServerIpv4  string
	ServerIpv6  string
	HostName    string
	Token       string
	PollAddress string
}

// UpdateInfo model
type UpdateInfo struct {
	Code         int64
	Sites        []DataSite
	Certificates []DataCertificate
	Services     []DataService
}

// DataSite model
type DataSite struct {
	ID     int64
	Domain string
	CerID  int64
	Config string
}

// DataService model
type DataService struct {
	Name    string
	Enable  string
	Disable string
	Status  string
}

// DataCertificate model
type DataCertificate struct {
	ID        int64
	Domain    string
	FullChain string
	Key       string
}

// StatusServer model
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

// StatusPercent model
type StatusPercent struct {
	CPU  float64
	Disk float64
	Mem  float64
	Swap float64
}

// CPUInfo model
type CPUInfo struct {
	ModelName string
	Cores     int32
}

// MemInfo model
type MemInfo struct {
	Total     uint64
	Used      uint64
	Available uint64
}

// SwapInfo model
type SwapInfo struct {
	Total     uint64
	Used      uint64
	Available uint64
}

// InterfaceInfo model
type InterfaceInfo struct {
	Addrs    []string
	ByteSent uint64
	ByteRecv uint64
}

// Webreq model
type Webreq struct {
	Code int64  `json:Code`
	Info string `json:Info`
}
