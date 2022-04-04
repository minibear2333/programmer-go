package model

import (
	"context"
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/zeromicro/go-zero/core/stores/mongo"
)

type UserModel interface {
	Insert(ctx context.Context, data *User) error
	FindOne(ctx context.Context, id string) (*User, error)
	FindBySearch(ctx context.Context, search string, pageNo int, pageSize int) (*[]User, error)
	FindUsersBySearchAndIds(ctx context.Context, search string, ids []bson.ObjectId, pageNo int, pageSize int) (*[]User, error)
	Update(ctx context.Context, data *User) error
	UpdateFields(ctx context.Context, id string, data *map[string]interface{}) error
	AddUserToSetByID(ctx context.Context, id string, key string, valueID string) error
	Delete(ctx context.Context, id string) error
	FindOneByOpenId(ctx context.Context, openId string) (*User, error)
}

type defaultUserModel struct {
	*mongo.Model
}

func NewUserModel(url, collection string) UserModel {
	return &defaultUserModel{
		Model: mongo.MustNewModel(url, collection),
	}
}

func (m *defaultUserModel) Insert(ctx context.Context, data *User) error {
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

func (m *defaultUserModel) FindOne(ctx context.Context, id string) (*User, error) {
	if !bson.IsObjectIdHex(id) {
		return nil, ErrInvalidObjectId
	}

	session, err := m.TakeSession()
	if err != nil {
		return nil, err
	}

	defer m.PutSession(session)
	var data User

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
func (m *defaultUserModel) FindBySearch(ctx context.Context, search string, pageNo int, pageSize int) (*[]User, error) {
	session, err := m.TakeSession()
	if err != nil {
		return nil, err
	}

	defer m.PutSession(session)
	var data []User
	filter := bson.M{
		"title": bson.M{"$regex": bson.RegEx{
			Pattern: fmt.Sprintf("%s", search),
			Options: "im",
		}}}
	count, err := m.GetCollection(session).Find(filter).Count()
	if err != nil {
		return nil, err
	}
	if count < pageNo {
		return &data, nil
	}
	skipNum := (pageNo - 1) * pageSize
	err = m.GetCollection(session).Find(filter).Skip(skipNum).Limit(pageSize).All(&data)
	switch err {
	case nil:
		return &data, nil
	case mongo.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) FindUsersBySearchAndIds(ctx context.Context, search string, oIDs []bson.ObjectId, pageNo int, pageSize int) (*[]User, error) {
	session, err := m.TakeSession()
	if err != nil {
		return nil, err
	}

	defer m.PutSession(session)
	var data []User
	filter := bson.M{
		"name": bson.M{"$regex": bson.RegEx{
			Pattern: fmt.Sprintf("%s", search),
			Options: "im",
		}},
		"_id": bson.M{
			"$in": oIDs,
		}}
	count, err := m.GetCollection(session).Find(filter).Count()
	if err != nil {
		return nil, err
	}
	if count < pageNo {
		return &data, nil
	}
	skipNum := (pageNo - 1) * pageSize
	err = m.GetCollection(session).Find(filter).Skip(skipNum).Limit(pageSize).All(&data)
	switch err {
	case nil:
		return &data, nil
	case mongo.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) Update(ctx context.Context, data *User) error {
	session, err := m.TakeSession()
	if err != nil {
		return err
	}

	defer m.PutSession(session)

	return m.GetCollection(session).UpdateId(data.ID, data)
}

func (m *defaultUserModel) Delete(ctx context.Context, id string) error {
	session, err := m.TakeSession()
	if err != nil {
		return err
	}

	defer m.PutSession(session)

	return m.GetCollection(session).RemoveId(bson.ObjectIdHex(id))
}

func (m *defaultUserModel) FindOneByOpenId(ctx context.Context, openId string) (*User, error) {
	session, err := m.TakeSession()
	if err != nil {
		return nil, err
	}

	defer m.PutSession(session)
	var data User

	filter := bson.D{{"open_id", openId}}
	err = m.GetCollection(session).Find(filter).One(&data)
	switch err {
	case nil:
		return &data, nil
	case mongo.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) UpdateFields(ctx context.Context, id string, data *map[string]interface{}) error {
	session, err := m.TakeSession()
	if err != nil {
		return err
	}

	defer m.PutSession(session)

	return m.GetCollection(session).Update(bson.M{"_id": bson.ObjectIdHex(id)},
		bson.M{"$set": data})
}

func (m *defaultUserModel) AddUserToSetByID(ctx context.Context, id string, key string, valueID string) error {
	if !bson.IsObjectIdHex(id) || !bson.IsObjectIdHex(valueID) {
		return ErrInvalidObjectId
	}

	session, err := m.TakeSession()
	if err != nil {
		return err
	}

	defer m.PutSession(session)

	return m.GetCollection(session).Update(bson.M{"_id": bson.ObjectIdHex(id)},
		bson.M{"$addToSet": bson.M{
			key: bson.ObjectIdHex(valueID),
		}})
}
