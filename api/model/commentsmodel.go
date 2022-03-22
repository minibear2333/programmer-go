package model

import (
	"context"

	"github.com/globalsign/mgo/bson"
	"github.com/zeromicro/go-zero/core/stores/mongo"
)

type CommentsModel interface {
	Insert(ctx context.Context, data *Comments) error
	FindOne(ctx context.Context, id string) (*Comments, error)
	Update(ctx context.Context, data *Comments) error
	Delete(ctx context.Context, id string) error
}

type defaultCommentsModel struct {
	*mongo.Model
}

func NewCommentsModel(url, collection string) CommentsModel {
	return &defaultCommentsModel{
		Model: mongo.MustNewModel(url, collection),
	}
}

func (m *defaultCommentsModel) Insert(ctx context.Context, data *Comments) error {
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

func (m *defaultCommentsModel) FindOne(ctx context.Context, id string) (*Comments, error) {
	if !bson.IsObjectIdHex(id) {
		return nil, ErrInvalidObjectId
	}

	session, err := m.TakeSession()
	if err != nil {
		return nil, err
	}

	defer m.PutSession(session)
	var data Comments

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

func (m *defaultCommentsModel) Update(ctx context.Context, data *Comments) error {
	session, err := m.TakeSession()
	if err != nil {
		return err
	}

	defer m.PutSession(session)

	return m.GetCollection(session).UpdateId(data.ID, data)
}

func (m *defaultCommentsModel) Delete(ctx context.Context, id string) error {
	session, err := m.TakeSession()
	if err != nil {
		return err
	}

	defer m.PutSession(session)

	return m.GetCollection(session).RemoveId(bson.ObjectIdHex(id))
}
