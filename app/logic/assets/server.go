package assets

import (
	"ControlCenter/model/mongoModel"
	"ControlCenter/pkg/utils"
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type Server struct {
	ctx    context.Context
	ID     string
	UserID string
}

func NewAssetsServer(ctx context.Context, id, userID string) *Server {
	return &Server{
		ctx:    ctx,
		ID:     id,
		UserID: userID,
	}
}

var _ Assets = (*Server)(nil)

func (s *Server) Get() (mongoModel.Model, error) {
	if !s.checkAuthority(mongoModel.AuthorityTypeRead) {
		return nil, errors.New("authority error")
	}
	var svr mongoModel.ModelServer
	err := svr.DB().FindOne(s.ctx, bson.M{
		"_id": s.ID,
	}).Decode(&svr)
	if err != nil {
		return nil, err
	}
	return &svr, nil
}

func (s *Server) Add(assets mongoModel.Model) error {
	svr, ok := assets.(*mongoModel.ModelServer)
	if !ok {
		return errors.New("cover error")
	}
	s.ID = svr.ID
	fmt.Println(s.ID, svr.ID)
	_, err := svr.DB().InsertOne(s.ctx, &svr)
	if err != nil {
		return err
	}
	var asset = mongoModel.ModelAssets{
		ID:         svr.ID,
		AssetsType: mongoModel.AssetsTypeServer,
		RemarkName: svr.RemarkName,
		Owner:      s.UserID,
		Authority: []*mongoModel.Authority{
			{UserID: s.UserID, Type: mongoModel.AuthorityTypeWrite},
			{UserID: s.UserID, Type: mongoModel.AuthorityTypeRead},
		},
		CreateAt: time.Now().UnixNano() / 1e6,
	}
	_, err = asset.DB().InsertOne(s.ctx, &asset)
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) Edit(assets mongoModel.Model) error {
	if !s.checkAuthority(mongoModel.AuthorityTypeWrite) {
		return errors.New("authority error")
	}
	svr, ok := assets.(*mongoModel.ModelServer)
	if !ok {
		return errors.New("cover error")
	}
	_, err := svr.DB().UpdateOne(s.ctx, bson.M{
		"_id": s.ID,
	}, svr)
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) Remove() error {
	if !s.checkAuthority(mongoModel.AuthorityTypeWrite) {
		return errors.New("authority error")
	}
	var svr mongoModel.ModelServer
	_, err := svr.DB().DeleteOne(s.ctx, bson.M{
		"_id": utils.RandomString(),
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) checkAuthority(authorityType int) bool {
	var assets mongoModel.ModelAssets
	err := assets.DB().FindOne(s.ctx, bson.M{
		"_id": s.ID,
		"authority": bson.M{
			"$elemMatch": bson.M{
				"user_id": s.UserID,
				"type":    authorityType,
			},
		},
	}).Decode(&assets)
	if err != nil {
		return false
	}
	return true
}
