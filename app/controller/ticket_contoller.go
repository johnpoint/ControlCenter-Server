package controller

import (
	"ControlCenter/infra"
	"ControlCenter/model/mongomodel"
	"ControlCenter/pkg/errorhelper"
	"ControlCenter/pkg/log"
	"ControlCenter/pkg/utils"
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
	user, err := getUserInfoByUserID(c, userID)
	if err != nil {
		returnErrorMsg(c, infra.ErrNeedVerifyInfo)
		return
	}
	var reqData ListTicketReq
	err = c.BindJSON(&reqData)
	if err != nil {
		returnErrorMsg(c, errorhelper.WarpErr(infra.ReqParseError, err))
		return
	}
	var resp PaginationResp
	var respItems []*ListTicketItem
	var ticket mongomodel.ModelTicket
	var filter = make(bson.M)
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
	var filter = make(bson.M)

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

type CreateTicketReq struct {
	Title   string                 `json:"title"`
	Content string                 `json:"content"`
	Level   mongomodel.TicketLevel `json:"level"`
}

type CreateTicketResp struct {
	ID string `json:"id"`
}

func CreateTicket(c *gin.Context) {
	var reqData CreateTicketReq
	err := c.BindJSON(&reqData)
	if err != nil {
		returnErrorMsg(c, errorhelper.WarpErr(infra.ReqParseError, err))
		return
	}
	userID, exists := getUserIDFromContext(c)
	if !exists {
		returnErrorMsg(c, infra.ErrNeedVerifyInfo)
		return
	}
	var ticket = mongomodel.ModelTicket{
		ID:        utils.RandomString(),
		Title:     reqData.Title,
		CreatedAt: time.Now().UnixNano() / 1e6,
		Level:     reqData.Level,
		UserID:    userID,
		Record: []*mongomodel.TicketRecord{
			{Sender: userID, Content: reqData.Content, CreatedAt: time.Now().UnixNano() / 1e6},
		},
		Status: mongomodel.TicketStatusPending,
	}
	id, err := ticket.DB().InsertOne(c, &ticket)
	if err != nil {
		returnErrorMsg(c, errorhelper.WarpErr(infra.DataBaseError, err))
		return
	}
	var resp CreateTicketResp
	resp.ID = id.InsertedID.(string)
	returnSuccessMsg(c, "", resp)
}

type PostTicketReq struct {
	Content string `json:"content"`
}

func PostTicket(c *gin.Context) {
	var reqData CreateTicketReq
	err := c.BindJSON(&reqData)
	if err != nil {
		returnErrorMsg(c, errorhelper.WarpErr(infra.ReqParseError, err))
		return
	}
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
	var filter = bson.M{
		"_id": c.Param("uuid"),
	}
	var updateStatus mongomodel.TicketStatus
	switch user.Power {
	case mongomodel.UserPowerOperation:
		filter["op_id"] = userID
		updateStatus = mongomodel.TicketStatusWaitOwner
	case mongomodel.UserPowerAdmin:
		updateStatus = mongomodel.TicketStatusWaitOwner
	case mongomodel.UserPowerSystem:
		updateStatus = mongomodel.TicketStatusWaitOwner
	default:
		filter["user_id"] = userID
		updateStatus = mongomodel.TicketStatusWaitOP
	}
	var ticket mongomodel.ModelTicket
	_, err = ticket.DB().UpdateOne(c, filter, bson.M{
		"$addToSet": bson.M{
			"record": &mongomodel.TicketRecord{
				Sender:    userID,
				Content:   reqData.Content,
				CreatedAt: time.Now().UnixNano() / 1e6,
			},
		},
		"$set": bson.M{
			"updated_at": time.Now().UnixNano() / 1e6,
			"status":     updateStatus,
		},
	})
	if err != nil {
		returnErrorMsg(c, errorhelper.WarpErr(infra.DataBaseError, err))
		return
	}
	returnSuccessMsg(c, "", nil)
}

func GetTicket(c *gin.Context) {
	userID, exists := getUserIDFromContext(c)
	if !exists {
		returnErrorMsg(c, infra.ErrNeedVerifyInfo)
		return
	}
	user, err := getUserInfoByUserID(c, userID)
	if err != nil {
		returnErrorMsg(c, infra.ErrNeedVerifyInfo)
		return
	}
	var ticket mongomodel.ModelTicket
	var filter = make(bson.M)
	filter["_id"] = c.Param("uuid")
	switch user.Power {
	case mongomodel.UserPowerSystem:
	case mongomodel.UserPowerAdmin:
	case mongomodel.UserPowerOperation:
		filter["op_id"] = userID
	default:
		filter["user_id"] = userID
	}
	ticket.DB().FindOne(c, bson.M{
		"_id": c.Param("uuid"),
	}).Decode(&ticket)
	if len(ticket.ID) == 0 {
		returnErrorMsg(c, infra.ErrNoPermission)
		return
	}
	var senderIDs = make([]string, 0)
	var nicknameMap = make(map[string]string)
	for i := range ticket.Record {
		senderIDs = append(senderIDs, ticket.Record[i].Sender)
	}
	if len(senderIDs) != 0 {
		var u mongomodel.ModelUser
		cur, err := u.DB().Find(c, bson.M{
			"_id": bson.M{
				"$in": senderIDs,
			},
		})
		if err != nil {
			log.Error("GetTicket", log.String("info", err.Error()))
		}
		defer cur.Close(c)
		for cur.Next(c) {
			err := cur.Decode(&u)
			if err != nil {
				continue
			}
			nicknameMap[u.ID] = u.Nickname
		}
	}
	for i := range ticket.Record {
		ticket.Record[i].IsSelf = ticket.Record[i].Sender == userID
		nickname, has := nicknameMap[ticket.Record[i].Sender]
		if has {
			ticket.Record[i].Sender = nickname
		}
	}

	returnSuccessMsg(c, "", ticket)
}
