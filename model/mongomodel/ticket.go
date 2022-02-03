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
	ID        string          `bson:"_id"`
	Title     string          `bson:"title"`
	OpID      string          `bson:"op_id"`
	UserID    string          `bson:"user_id"`
	Level     TicketLevel     `bson:"level"`
	Status    TicketStatus    `bson:"status"`
	Record    []*TicketRecord `bson:"record"`
	CreatedAt int64           `bson:"created_at"`
	UpdatedAt int64           `bson:"updated_at"`
}

type TicketRecord struct {
	Sender    string `bson:"sender"`
	Content   string `bson:"content"`
	CreatedAt int64  `bson:"created_at"`
}

func (m *ModelTicket) CollectionName() string {
	return "ticket"
}

func (m *ModelTicket) DB() *mongo.Collection {
	return mongodao.Collection(m.CollectionName())
}
