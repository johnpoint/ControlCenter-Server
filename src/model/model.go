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
		Debug        bool
		RedisConfig  struct {
			Addr     string
			Password string
			DB       int
		}
	}

	// UpdateInfo model
	UpdateInfo struct {
		Code  int64
		Sites []struct {
			ID     int64
			Domain string
			CerID  int64
			Config string
		}
		Certificates []struct {
			ID        int64
			Domain    string
			FullChain string
			Key       string
		}
		Services struct {
			Name    string
			Enable  string
			Disable string
			Status  string
		}
		Dockers []DockerInfo
	}
)
