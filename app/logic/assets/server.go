package assets

import (
	"ControlCenter/model/mongomodel"
	"ControlCenter/pkg/utils"
	"context"
	"errors"
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

func (s *Server) Get() (mongomodel.Model, error) {
	if !s.checkAuthority(mongomodel.AuthorityTypeRead) {
		return nil, errors.New("authority error")
	}
	var svr mongomodel.ModelServer
	err := svr.DB().FindOne(s.ctx, bson.M{
		"_id": s.ID,
	}).Decode(&svr)
	if err != nil {
		return nil, err
	}
	return &svr, nil
}

func (s *Server) Add(assets mongomodel.Model) error {
	svr, ok := assets.(*mongomodel.ModelServer)
	if !ok {
		return errors.New("cover error")
	}
	s.ID = svr.ID
	_, err := svr.DB().InsertOne(s.ctx, &svr)
	if err != nil {
		return err
	}
	var asset = mongomodel.ModelAssets{
		ID:         svr.ID,
		AssetsType: mongomodel.AssetsTypeServer,
		RemarkName: svr.RemarkName,
		Owner:      s.UserID,
		Authority: []*mongomodel.Authority{
			{UserID: s.UserID, Type: mongomodel.AuthorityTypeWrite},
			{UserID: s.UserID, Type: mongomodel.AuthorityTypeRead},
		},
		CreateAt: time.Now().UnixNano() / 1e6,
	}
	_, err = asset.DB().InsertOne(s.ctx, &asset)
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) Edit(assets mongomodel.Model) error {
	if !s.checkAuthority(mongomodel.AuthorityTypeWrite) {
		return errors.New("authority error")
	}
	svr, ok := assets.(*mongomodel.ModelServer)
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
	if !s.checkAuthority(mongomodel.AuthorityTypeWrite) {
		return errors.New("authority error")
	}
	var svr mongomodel.ModelServer
	_, err := svr.DB().DeleteOne(s.ctx, bson.M{
		"_id": utils.RandomString(),
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) checkAuthority(authorityType int) bool {
	var assets mongomodel.ModelAssets
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
