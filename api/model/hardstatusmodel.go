package model

import (
	"context"

	"github.com/globalsign/mgo/bson"
	"github.com/zeromicro/go-zero/core/stores/mongo"
)

type HardStatusModel interface {
	Insert(ctx context.Context, data *HardStatus) error
	FindOne(ctx context.Context, id string) (*HardStatus, error)
	Update(ctx context.Context, data *HardStatus) error
	Delete(ctx context.Context, id string) error
}

type defaultHardStatusModel struct {
	*mongo.Model
}

func NewHardStatusModel(url, collection string) HardStatusModel {
	return &defaultHardStatusModel{
		Model: mongo.MustNewModel(url, collection),
	}
}

func (m *defaultHardStatusModel) Insert(ctx context.Context, data *HardStatus) error {
	if !data.ID.Valid() {
		data.ID = bson.NewObjectId()
	}

	session, err := m.TakeSession()
	if err != nil {
		return err
	}

	defer m.PutSession(session)
	return m.GetCollection(session).Insert(data)
}

func (m *defaultHardStatusModel) FindOne(ctx context.Context, id string) (*HardStatus, error) {
	if !bson.IsObjectIdHex(id) {
		return nil, ErrInvalidObjectId
	}

	session, err := m.TakeSession()
	if err != nil {
		return nil, err
	}

	defer m.PutSession(session)
	var data HardStatus

	err = m.GetCollection(session).FindId(bson.ObjectIdHex(id)).One(&data)
	switch err {
	case nil:
		return &data, nil
	case mongo.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultHardStatusModel) Update(ctx context.Context, data *HardStatus) error {
	session, err := m.TakeSession()
	if err != nil {
		return err
	}

	defer m.PutSession(session)

	return m.GetCollection(session).UpdateId(data.ID, data)
}

func (m *defaultHardStatusModel) Delete(ctx context.Context, id string) error {
	session, err := m.TakeSession()
	if err != nil {
		return err
	}

	defer m.PutSession(session)

	return m.GetCollection(session).RemoveId(bson.ObjectIdHex(id))
}
