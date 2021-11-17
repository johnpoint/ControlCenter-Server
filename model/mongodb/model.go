package mongodb

type Model interface {
	CollectionName() string
}

type DefaultModel struct{}

func (d *DefaultModel) CollectionName() string {
	return ""
}

// 检查是否实现接口
var _ Model = (*DefaultModel)(nil)

var _ Model = (*ModelUser)(nil)
var _ Model = (*ModelServer)(nil)
var _ Model = (*ModelAssets)(nil)
