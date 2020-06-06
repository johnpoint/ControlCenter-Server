package model

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
	}

	// SysConfig model of config

	// UpdateInfo model
	UpdateInfo struct {
		Code         int64
		Sites        []DataSite
		Certificates []DataCertificate
		Services     []DataService
		Dockers      []DockerInfo
	}

	// DataSite model
	DataSite struct {
		ID     int64
		Domain string
		CerID  int64
		Config string
	}

	// DataService model
	DataService struct {
		Name    string
		Enable  string
		Disable string
		Status  string
	}

	// DataCertificate model
	DataCertificate struct {
		ID        int64
		Domain    string
		FullChain string
		Key       string
	}
)
