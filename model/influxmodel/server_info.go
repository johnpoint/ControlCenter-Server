package influxmodel

type ModelServerInfo struct {
	CPU    float64 `json:"cpu"`
	Disk   float64 `json:"disk"`
	Memory float64 `json:"memory"`
	Swap   float64 `json:"swap"`
}

var _ Model = (*ModelServerInfo)(nil)

func (m *ModelServerInfo) BucketName() string {
	return "server_info"
}

func (m *ModelServerInfo) Measurement() string {
	return "server_performance"
}
