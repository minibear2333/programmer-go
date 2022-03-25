package model

import (
	"context"

	"github.com/globalsign/mgo/bson"
	"github.com/zeromicro/go-zero/core/stores/mongo"
)

type MessageCenterModel interface {
	Insert(ctx context.Context, data *MessageCenter) error
	FindOne(ctx context.Context, id string) (*MessageCenter, error)
	Update(ctx context.Context, data *MessageCenter) error
	Delete(ctx context.Context, id string) error
}

type defaultMessageCenterModel struct {
	*mongo.Model
}

func NewMessageCenterModel(url, collection string) MessageCenterModel {
	return &defaultMessageCenterModel{
		Model: mongo.MustNewModel(url, collection),
	}
}

func (m *defaultMessageCenterModel) Insert(ctx context.Context, data *MessageCenter) error {
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

func (m *defaultMessageCenterModel) FindOne(ctx context.Context, id string) (*MessageCenter, error) {
	if !bson.IsObjectIdHex(id) {
		return nil, ErrInvalidObjectId
	}

	session, err := m.TakeSession()
	if err != nil {
		return nil, err
	}

	defer m.PutSession(session)
	var data MessageCenter

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

func (m *defaultMessageCenterModel) Update(ctx context.Context, data *MessageCenter) error {
	session, err := m.TakeSession()
	if err != nil {
		return err
	}

	defer m.PutSession(session)

	return m.GetCollection(session).UpdateId(data.ID, data)
}

func (m *defaultMessageCenterModel) Delete(ctx context.Context, id string) error {
	session, err := m.TakeSession()
	if err != nil {
		return err
	}

	defer m.PutSession(session)

	return m.GetCollection(session).RemoveId(bson.ObjectIdHex(id))
}
