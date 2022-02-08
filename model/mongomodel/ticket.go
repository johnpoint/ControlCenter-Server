package mongomodel

import (
	"ControlCenter/dao/mongodao"
	"go.mongodb.org/mongo-driver/mongo"
)

type TicketLevel = int32

const (
	TicketLevelLow TicketLevel = iota + 1
	TicketLevelMiddle
	TicketLevelHigh
)

type TicketStatus = int32

const (
	TicketStatusPending TicketStatus = iota + 1
	TicketStatusClosed
	TicketStatusWaitOP
	TicketStatusWaitOwner
)

type ModelTicket struct {
	ID        string          `bson:"_id" json:"id"`
	Title     string          `bson:"title" json:"title"`
	OpID      string          `bson:"op_id" json:"op_id"`
	UserID    string          `bson:"user_id" json:"user_id"`
	Level     TicketLevel     `bson:"level" json:"level"`
	Status    TicketStatus    `bson:"status" json:"status"`
	Record    []*TicketRecord `bson:"record" json:"record"`
	CreatedAt int64           `bson:"created_at" json:"created_at"`
	UpdatedAt int64           `bson:"updated_at" json:"updated_at"`
}

type TicketRecord struct {
	Sender    string `bson:"sender" json:"sender"`
	Content   string `bson:"content" json:"content"`
	CreatedAt int64  `bson:"created_at" json:"created_at"`
	IsSelf    bool   `json:"is_self" bson:"-"`
}

func (m *ModelTicket) CollectionName() string {
	return "ticket"
}

func (m *ModelTicket) DB() *mongo.Collection {
	return mongodao.Collection(m.CollectionName())
}
