package influxDB

type Config struct {
	Address string `json:"address"`
	Token   string `json:"token"`
	Org     string `json:"org"`
}
