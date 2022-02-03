package controller

import (
	"ControlCenter/infra"
	"ControlCenter/model/mongomodel"
	"ControlCenter/pkg/errorhelper"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type ListTicketReq struct {
	Status mongomodel.TicketStatus `json:"status"`
	Level  mongomodel.TicketLevel  `json:"level"`
	PaginationReq
}

type ListTicketItem struct {
	ID        string                  `json:"id"`
	Title     string                  `json:"title"`
	Status    mongomodel.TicketStatus `json:"status"`
	Level     mongomodel.TicketLevel  `json:"level"`
	CreatedAt int64                   `json:"created_at"`
}

func ListTicket(c *gin.Context) {
	userID, exists := getUserIDFromContext(c)
	if !exists {
		returnErrorMsg(c, infra.ErrNeedVerifyInfo)
		return
	}
	var user mongomodel.ModelUser
	_ = user.DB().FindOne(c, bson.M{"_id": userID}).Decode(&user)
	if len(user.ID) == 0 {
		returnErrorMsg(c, infra.ErrNeedVerifyInfo)
		return
	}
	var reqData ListTicketReq
	err := c.BindJSON(&reqData)
	if err != nil {
		returnErrorMsg(c, errorhelper.WarpErr(infra.ReqParseError, err))
		return
	}
	var resp PaginationResp
	var respItems []*ListTicketItem
	var ticket mongomodel.ModelTicket
	var filter bson.M
	var opts options.FindOptions

	if reqData.PageSize > 0 && reqData.Page > 0 {
		opts.SetLimit(reqData.PageSize)
		opts.SetSkip((reqData.Page - 1) * reqData.PageSize)
		resp.PerPage = reqData.PageSize
		resp.Page = reqData.Page
	} else {
		opts.SetLimit(10)
	}
	if reqData.Level != 0 {
		filter["level"] = reqData.Level
	}
	if reqData.Status != 0 {
		filter["status"] = reqData.Status
	}
	opts.SetSort(bson.D{{Key: "created_at", Value: -1}})
	switch user.Power {
	case mongomodel.UserPowerAdmin:
	case mongomodel.UserPowerSystem:
	case mongomodel.UserPowerOperation:
		filter["op_id"] = userID
	default:
		filter["user_id"] = userID
	}
	cur, err := ticket.DB().Find(c, filter, &opts)
	if err != nil {
		returnErrorMsg(c, errorhelper.WarpErr(infra.DataBaseError, err))
		return
	}
	var tickets = make([]*mongomodel.ModelTicket, 0)
	err = cur.All(c, &tickets)
	if err != nil {
		returnErrorMsg(c, errorhelper.WarpErr(infra.DataBaseError, err))
		return
	}
	resp.Total, _ = ticket.DB().CountDocuments(c, filter)
	for i := range tickets {
		respItems = append(respItems, &ListTicketItem{
			ID:        tickets[i].ID,
			Title:     tickets[i].Title,
			Status:    tickets[i].Status,
			Level:     tickets[i].Level,
			CreatedAt: tickets[i].CreatedAt,
		})
	}
	resp.Data = respItems
	returnSuccessMsg(c, "", resp)
}

func CloseTicket(c *gin.Context) {
	userID, exists := getUserIDFromContext(c)
	if !exists {
		returnErrorMsg(c, infra.ErrNeedVerifyInfo)
		return
	}
	var user mongomodel.ModelUser
	_ = user.DB().FindOne(c, bson.M{"_id": userID}).Decode(&user)
	if len(user.ID) == 0 {
		returnErrorMsg(c, infra.ErrNeedVerifyInfo)
		return
	}
	uuid := c.Param("uuid")
	var ticket mongomodel.ModelTicket
	var filter bson.M

	filter["_id"] = uuid
	switch user.Power {
	case mongomodel.UserPowerAdmin:
	case mongomodel.UserPowerSystem:
	case mongomodel.UserPowerOperation:
		filter["op_id"] = userID
	default:
		filter["user_id"] = userID
	}
	_ = ticket.DB().FindOne(c, filter).Decode(&ticket)
	if len(ticket.ID) == 0 {
		returnErrorMsg(c, infra.ErrNoPermission)
		return
	}
	_, err := ticket.DB().UpdateOne(c, filter, bson.M{
		"$set": bson.M{
			"status":     mongomodel.TicketStatusClosed,
			"updated_at": time.Now().UnixNano() / 1e6,
		},
	})
	if err != nil {
		returnErrorMsg(c, errorhelper.WarpErr(infra.DataBaseError, err))
		return
	}
	returnSuccessMsg(c, "", nil)
}
