package mongodb

type ModelServer struct {
	ID string `json:"_id" bson:"_id"`
}

func (m *ModelServer) CollectionName() string {
	return "server"
}
