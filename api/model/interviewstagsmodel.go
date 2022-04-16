package model

import (
	"context"
	"github.com/globalsign/mgo/bson"
	"github.com/zeromicro/go-zero/core/stores/mongo"
	"strings"
)

type InterviewsTagsModel interface {
	Insert(ctx context.Context, data *InterviewsTags) error
	FindOne(ctx context.Context, id string) (*InterviewsTags, error)
	CheckTag(ctx context.Context, tags []string) (bool, error)
	FindAll(ctx context.Context) (*[]InterviewsTags, error)
	Update(ctx context.Context, data *InterviewsTags) error
	Delete(ctx context.Context, id string) error
}

type defaultInterviewsTagsModel struct {
	*mongo.Model
}

func NewInterviewsTagsModel(url, collection string) InterviewsTagsModel {
	return &defaultInterviewsTagsModel{
		Model: mongo.MustNewModel(url, collection),
	}
}

func (m *defaultInterviewsTagsModel) Insert(ctx context.Context, data *InterviewsTags) error {
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

func (m *defaultInterviewsTagsModel) FindOne(ctx context.Context, id string) (*InterviewsTags, error) {
	if !bson.IsObjectIdHex(id) {
		return nil, ErrInvalidObjectId
	}

	session, err := m.TakeSession()
	if err != nil {
		return nil, err
	}

	defer m.PutSession(session)
	var data InterviewsTags

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

func (m *defaultInterviewsTagsModel) CheckTag(ctx context.Context, tags []string) (bool, error) {

	session, err := m.TakeSession()
	if err != nil {
		return false, err
	}

	defer m.PutSession(session)
	var data InterviewsTags
	for i := range tags {
		tagArr := strings.Split(tags[i], "-")
		if len(tagArr) != 2 {
			return false, nil
		}
		query := bson.M{"name": tagArr[0], "sub_tags": bson.M{"$elemMatch": bson.M{"$eq": tagArr[1]}}}
		err = m.GetCollection(session).Find(query).One(&data)
		switch err {
		case mongo.ErrNotFound:
			return false, ErrNotFound
		default:
			if err!=nil{
				return false, err
			}
		}
	}
	return true, nil
}

func (m *defaultInterviewsTagsModel) FindAll(ctx context.Context) (*[]InterviewsTags, error) {
	session, err := m.TakeSession()
	if err != nil {
		return nil, err
	}

	defer m.PutSession(session)
	var data []InterviewsTags
	err = m.GetCollection(session).Find(nil).All(&data)
	switch err {
	case nil:
		return &data, nil
	case mongo.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultInterviewsTagsModel) Update(ctx context.Context, data *InterviewsTags) error {
	session, err := m.TakeSession()
	if err != nil {
		return err
	}

	defer m.PutSession(session)

	return m.GetCollection(session).UpdateId(data.ID, data)
}

func (m *defaultInterviewsTagsModel) Delete(ctx context.Context, id string) error {
	session, err := m.TakeSession()
	if err != nil {
		return err
	}

	defer m.PutSession(session)

	return m.GetCollection(session).RemoveId(bson.ObjectIdHex(id))
}
