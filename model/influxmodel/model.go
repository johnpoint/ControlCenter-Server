package influxmodel

type Model interface {
	BucketName() string
	Measurement() string
}

type DefaultModel struct{}

func (d *DefaultModel) BucketName() string {
	return ""
}

func (d *DefaultModel) Measurement() string {
	return ""
}

// 检查是否实现接口
var _ Model = (*DefaultModel)(nil)
