package mongodb

type AssetsType int32

type Authority struct {
	UserID string `json:"user_id" bson:"user_id"`
	Write  bool   `json:"write" bson:"write"`
	Read   bool   `json:"read" bson:"read"`
}

type Assets struct {
	ID         string      `json:"id" bson:"_id"`
	AssetsType AssetsType  `json:"assets_type" bson:"assets_type"`
	Owner      string      `json:"owner" bson:"owner"`
	Authority  []Authority `json:"authority" bson:"authority"`
}

func (a *Assets) CollectionName() string {
	return "assets"
}
