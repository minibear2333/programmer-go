package model

import (
	"context"

	"github.com/globalsign/mgo/bson"
	"github.com/zeromicro/go-zero/core/stores/mongo"
)

type MessageConfigModel interface {
	Insert(ctx context.Context, data *MessageConfig) error
	FindOne(ctx context.Context, id string) (*MessageConfig, error)
	FindOneByUserID(ctx context.Context, userID string) (*MessageConfig, error)
	Update(ctx context.Context, data *MessageConfig) error
	UpdateByUserID(ctx context.Context, data *MessageConfig) error
	Delete(ctx context.Context, id string) error
}

type defaultMessageConfigModel struct {
	*mongo.Model
}

func NewMessageConfigModel(url, collection string) MessageConfigModel {
	return &defaultMessageConfigModel{
		Model: mongo.MustNewModel(url, collection),
	}
}

func (m *defaultMessageConfigModel) Insert(ctx context.Context, data *MessageConfig) error {
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

func (m *defaultMessageConfigModel) FindOne(ctx context.Context, id string) (*MessageConfig, error) {
	if !bson.IsObjectIdHex(id) {
		return nil, ErrInvalidObjectId
	}

	session, err := m.TakeSession()
	if err != nil {
		return nil, err
	}

	defer m.PutSession(session)
	var data MessageConfig

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

func (m *defaultMessageConfigModel) FindOneByUserID(ctx context.Context, userID string) (*MessageConfig, error) {
	if !bson.IsObjectIdHex(userID) {
		return nil, ErrInvalidObjectId
	}

	session, err := m.TakeSession()
	if err != nil {
		return nil, err
	}

	defer m.PutSession(session)
	var data MessageConfig

	err = m.GetCollection(session).Find(bson.M{"user_id": bson.ObjectIdHex(userID)}).One(&data)
	switch err {
	case nil:
		return &data, nil
	case mongo.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultMessageConfigModel) Update(ctx context.Context, data *MessageConfig) error {
	session, err := m.TakeSession()
	if err != nil {
		return err
	}

	defer m.PutSession(session)

	return m.GetCollection(session).UpdateId(data.ID, data)
}

func (m *defaultMessageConfigModel) UpdateByUserID(ctx context.Context, data *MessageConfig) error {
	session, err := m.TakeSession()
	if err != nil {
		return err
	}

	defer m.PutSession(session)

	return m.GetCollection(session).Update(bson.M{"user_id": bson.ObjectIdHex(data.UserID.Hex())}, data)
}


func (m *defaultMessageConfigModel) Delete(ctx context.Context, id string) error {
	session, err := m.TakeSession()
	if err != nil {
		return err
	}

	defer m.PutSession(session)

	return m.GetCollection(session).RemoveId(bson.ObjectIdHex(id))
}
