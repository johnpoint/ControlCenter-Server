package mongodb

type Model interface {
	CollectionName() string
}

type DefaultModel struct{}

func (d *DefaultModel) CollectionName() string {
	return ""
}
