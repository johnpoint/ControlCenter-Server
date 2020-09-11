package model

import "github.com/shirou/gopsutil/load"

type (
	// Callback send some info to client
	Callback struct {
		Code int64
		Info string
	}

	// Config model of config
	Config struct {
		AllowAddress []string
		ListenPort   string
		TLS          bool
		CERTPath     string
		KEYPath      string
		Salt         string
		Database     string
		Debug        bool
		RedisConfig  struct {
			Enable   bool
			Addr     string
			Password string
			DB       int
		}
	}

	// UpdateInfo model
	UpdateInfo struct {
		Code         int64
		Certificates []DataCertificate
		ConfFile     []File
	}

	File struct {
		Name  string
		Path  string
		Value string
	}

	DataSite struct {
		ID     int64
		Domain string
		CerID  int64
		Config string
	}
	DataCertificate struct {
		ID        int64
		Domain    string
		FullChain string
		Key       string
	}
	// StatusServer model
	StatusServer struct {
		Version  string
		Percent  StatusPercent
		Load     *load.AvgStat
		BootTime uint64
		Uptime   uint64
	}

	// StatusPercent model
	StatusPercent struct {
		CPU  float64
		Disk float64
		Mem  float64
		Swap float64
	}
)

type ReSetPassword struct {
	Oldpass string `json:"oldpass" xml:"oldpass" form:"oldpass" query:"oldpass"`
	Newpass string `json:"newpass" xml:"newpass" form:"newpass" query:"newpass"`
}
