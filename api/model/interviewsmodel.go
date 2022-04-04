package model

import (
	"context"
	"fmt"

	"github.com/minibear2333/programmer-go/api/internal/types"

	"github.com/globalsign/mgo/bson"
	"github.com/zeromicro/go-zero/core/stores/mongo"
)

type InterviewsModel interface {
	Insert(ctx context.Context, data *Interviews) error
	FindOne(ctx context.Context, id string) (*Interviews, error)
	CountHardStatus(ctx context.Context, result *[]CountResult) error
	FindByTagsAndSearchWord(ctx context.Context, oIDs []bson.ObjectId, tags []string, search string, page types.CommonPage) (*[]Interviews, error)
	Update(ctx context.Context, data *Interviews) error
	UpdateFields(ctx context.Context, id string, data *map[string]interface{}) error
	Delete(ctx context.Context, id string) error
}

type defaultInterviewsModel struct {
	*mongo.Model
}

func NewInterviewsModel(url, collection string) InterviewsModel {
	return &defaultInterviewsModel{
		Model: mongo.MustNewModel(url, collection),
	}
}

func (m *defaultInterviewsModel) Insert(ctx context.Context, data *Interviews) error {
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

func (m *defaultInterviewsModel) FindOne(ctx context.Context, id string) (*Interviews, error) {
	if !bson.IsObjectIdHex(id) {
		return nil, ErrInvalidObjectId
	}

	session, err := m.TakeSession()
	if err != nil {
		return nil, err
	}

	defer m.PutSession(session)
	var data Interviews

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

func (m *defaultInterviewsModel) CountHardStatus(ctx context.Context, result *[]CountResult) error {
	session, err := m.TakeSession()
	if err != nil {
		return err
	}

	defer m.PutSession(session)
	err = m.GetCollection(session).Pipe([]bson.M{{"$group": bson.M{"_id": "$hard_status", "count": bson.M{"$count": bson.M{}}}}}).All(result)
	return err
}

func (m *defaultInterviewsModel) FindByTagsAndSearchWord(ctx context.Context, oIDs []bson.ObjectId, tags []string, search string, page types.CommonPage) (*[]Interviews, error) {
	session, err := m.TakeSession()
	if err != nil {
		return nil, err
	}

	defer m.PutSession(session)
	var data []Interviews
	filter := bson.M{}
	if tags != nil && len(tags) > 0 {
		filter["tags"] = tags
	}
	if oIDs != nil && len(oIDs) > 0 {
		filter["_id"] = bson.M{
			"$in": oIDs,
		}
	}
	if search != ""{
		filter["title"]= bson.M{"$regex": bson.RegEx{
			Pattern: fmt.Sprintf("%s", search),
			Options: "im",
		}}
	}
	count, err := m.GetCollection(session).Find(filter).Count()
	if err != nil {
		return nil, err
	}

	if count < page.PageNo {
		return &data, nil
	}

	skipNum := (page.PageNo - 1) * page.PageSize
	err = m.GetCollection(session).Find(filter).Skip(skipNum).Limit(page.PageSize).All(&data)
	switch err {
	case nil:
		return &data, nil
	case mongo.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultInterviewsModel) Update(ctx context.Context, data *Interviews) error {
	session, err := m.TakeSession()
	if err != nil {
		return err
	}

	defer m.PutSession(session)

	return m.GetCollection(session).UpdateId(data.ID, data)
}

func (m *defaultInterviewsModel) Delete(ctx context.Context, id string) error {
	session, err := m.TakeSession()
	if err != nil {
		return err
	}

	defer m.PutSession(session)

	return m.GetCollection(session).RemoveId(bson.ObjectIdHex(id))
}

func (m *defaultInterviewsModel) UpdateFields(ctx context.Context, id string, data *map[string]interface{}) error {
	session, err := m.TakeSession()
	if err != nil {
		return err
	}

	defer m.PutSession(session)

	return m.GetCollection(session).Update(bson.M{"_id": bson.ObjectIdHex(id)},
		bson.M{"$set": data})
}
