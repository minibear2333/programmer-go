package model

//go:generate goctl model mongo -t Comments -t HardStatus -t Interviews -t InterviewsTags -t User
import (
	"github.com/globalsign/mgo/bson"
	"time"
)

type Comments struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	ArticleID   bson.ObjectId `json:"article_id"`
	Avatar      string        `json:"avatar"`
	Content     string        `json:"content"`
	CreatedTime time.Time     `json:"created_time"`
	Good        int64         `json:"good"`
	Name        string        `json:"name"`
	Reply       []struct {
		Avatar      string    `json:"avatar"`
		Content     string    `json:"content"`
		CreatedTime time.Time `json:"created_time"`
		Name        string    `json:"name"`
		UserID      struct {
			Oid string `json:"$oid"`
		} `json:"user_id"`
	} `json:"reply"`
	UserID bson.ObjectId `json:"user_id"`
}

type HardStatus struct {
	ID     bson.ObjectId `bson:"_id"`
	Easy   int64         `json:"easy"`
	Hard   int64         `json:"hard"`
	Medium int64         `json:"medium"`
	Type   string        `json:"type"`
}

type Interviews struct {
	ID     bson.ObjectId `bson:"_id"`
	Author struct {
		ID struct {
			Oid string `json:"$oid"`
		} `json:"_id"`
		Name string `json:"name"`
	} `json:"author"`
	Bad      int64 `json:"bad"`
	ClickNum int64 `json:"click_num"`
	Comments []struct {
		ID struct {
			Oid string `json:"$oid"`
		} `json:"_id"`
	} `json:"comments"`
	CreatedTime time.Time `json:"created_time"`
	Good        int64     `json:"good"`
	HardStatus  string    `json:"hard_status"`
	HotNum      int64     `json:"hot_num"`
	StarNum     int64     `json:"star_num"`
	Summary     string    `json:"summary"`
	Tags        []string  `json:"tags"`
	Title       string    `json:"title"`
	UpdatedTime time.Time `json:"updated_time"`
}

type InterviewsTags struct {
	ID      bson.ObjectId `bson:"_id" json:"_id"`
	Name    string        `bson:"name" json:"name"`
	SubTags []string      `bson:"sub_tags" json:"sub_tags"`
}

type User struct {
	ID        bson.ObjectId `bson:"_id"`
	Avatar    string        `bson:"avatar" json:"avatar"`
	Birthday  time.Time     `bson:"birthday" json:"birthday"`
	Blog      string        `bson:"blog" json:"blog"`
	City      string        `bson:"city" json:"city"`
	Email     string        `bson:"email" json:"email"`
	Followers []struct {
		ID bson.ObjectId `bson:"_id" json:"_id"`
	} `json:"followers"`
	Following []struct {
		ID bson.ObjectId `bson:"_id" json:"_id"`
	} `json:"following"`
	Interviews struct {
		HardStatus struct {
			Easy   int64 `bson:"easy" json:"easy"`
			Hard   int64 `bson:"hard" json:"hard"`
			Medium int64 `bson:"medium" json:"medium"`
		} `bson:"hard_status" json:"hard_status"`
	} `bson:"interviews" json:"interviews"`
	Name     string   `bson:"name" json:"name"`
	Phone    int64    `bson:"phone" json:"phone"`
	RealName string   `bson:"real_name" json:"real_name"`
	Skills   []string `bson:"skills" json:"skills"`
	Star     []struct {
		Data []struct {
			ID          bson.ObjectId `bson:"_id" json:"_id"`
			Title       string        `bson:"title" json:"title"`
			UpdatedTime time.Time     `bson:"updated_time" json:"updated_time"`
		} `bson:"data" json:"data"`
		Name string `bson:"name" json:"name"`
		Type string `bson:"type" json:"type"`
	} `bson:"star" json:"star"`
	Summary string `bson:"summary" json:"summary"`
}
