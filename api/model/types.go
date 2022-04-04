package model

//go:generate goctl model mongo -t Comments -t HardStatus -t Interviews -t InterviewsTags -t User -t MessageCenter -t MessageConfig
import (
	"github.com/globalsign/mgo/bson"
	"time"
)

type Comments struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	ArticleID   bson.ObjectId `bson:"article_id" json:"article_id"`
	Avatar      string        `bson:"avatar" json:"avatar"`
	Content     string        `bson:"content" json:"content"`
	CreatedTime time.Time     `bson:"created_time" json:"created_time"`
	Good        int64         `bson:"good" json:"good"`
	Name        string        `bson:"name" json:"name"`
	Reply       []struct {
		Avatar      string        `bson:"avatar" json:"avatar"`
		Content     string        `bson:"content" json:"content"`
		CreatedTime time.Time     `bson:"created_time" json:"created_time"`
		Name        string        `bson:"name" json:"name"`
		UserID      bson.ObjectId `bson:"user_id" json:"user_id"`
	} `bson:"reply" json:"reply"`
	UserID bson.ObjectId `bson:"user_id" json:"user_id"`
}

type HardStatus struct {
	ID     bson.ObjectId `bson:"_id" json:"_id"`
	Easy   int64         `bson:"easy" json:"easy"`
	Hard   int64         `bson:"hard" json:"hard"`
	Medium int64         `bson:"medium" json:"medium"`
	Type   string        `bson:"type" json:"type"`
}
type Author struct {
	ID   bson.ObjectId `bson:"_id" json:"_id"`
	Name string        `bson:"name" json:"name"`
}
type CommentsID struct {
	ID bson.ObjectId `bson:"_id" json:"_id"`
}
type Interviews struct {
	ID          bson.ObjectId `bson:"_id" json:"_id"`
	Author      Author        `bson:"author" json:"author"`
	Content     string        `bson:"content" json:"content"`
	Bad         int64         `bson:"bad" json:"bad"`
	ClickNum    int64         `bson:"click_num" json:"click_num"`
	Comments    []CommentsID  `bson:"comments" json:"comments"`
	CreatedTime time.Time     `bson:"created_time" json:"created_time"`
	Good        int64         `bson:"good" json:"good"`
	HardStatus  string        `bson:"hard_status" json:"hard_status"`
	HotNum      int64         `bson:"hot_num" json:"hot_num"`
	StarNum     int64         `bson:"star_num" json:"star_num"`
	Summary     string        `bson:"summary" json:"summary"`
	Tags        []string      `bson:"tags" json:"tags"`
	Title       string        `bson:"title" json:"title"`
	UpdatedTime time.Time     `bson:"updated_time" json:"updated_time"`
}
type CountResult struct {
	ID    string `bson:"_id"`
	Count int64  `bson:"count"`
}

type InterviewsTags struct {
	ID      bson.ObjectId `bson:"_id" json:"_id"`
	Name    string        `bson:"name" json:"name"`
	SubTags []string      `bson:"sub_tags" json:"sub_tags"`
}

type User struct {
	ID        bson.ObjectId `bson:"_id" json:"_id"`
	OpenId    string        `bson:"open_id" json:"open_id"`
	Avatar    string        `bson:"avatar" json:"avatar"`
	Birthday  time.Time     `bson:"birthday" json:"birthday"`
	Blog      string        `bson:"blog" json:"blog"`
	City      string        `bson:"city" json:"city"`
	Email     string        `bson:"email" json:"email"`
	Followers []struct {
		ID bson.ObjectId `bson:"_id" json:"_id"`
	} `bson:"followers" json:"followers"`
	Following []struct {
		ID bson.ObjectId `bson:"_id" json:"_id"`
	} `bson:"following" json:"following"`
	InterviewsStatus map[string]int64 `bson:"interviews_status" json:"interviews_status"`
	Name             string           `bson:"name" json:"name"`
	Phone            int64            `bson:"phone" json:"phone"`
	RealName         string           `bson:"real_name" json:"real_name"`
	Skills           []string         `bson:"skills" json:"skills"`
	StarInterviews   []struct {
		ID bson.ObjectId `bson:"_id" json:"_id"`
	} `bson:"star_interviews" json:"star_interviews"`
	Summary string `bson:"summary" json:"summary"`
}

type MessageCenter struct {
	ID          bson.ObjectId `bson:"_id" json:"_id"`
	Content     string        `bson:"content" json:"content"`
	CreatedTime time.Time     `bson:"created_time" json:"created_time"`
	Title       string        `bson:"title" json:"title"`
	Type        int64         `bson:"type" json:"type"`
	User        struct {
		ID     bson.ObjectId `bson:"_id" json:"_id"`
		Action string        `bson:"action" json:"action"`
		Avatar string        `bson:"avatar" json:"avatar"`
		Name   string        `bson:"name" json:"name"`
	} `bson:"user" json:"user"`
}

type MessageConfig struct {
	ID          bson.ObjectId `bson:"_id" json:"_id"`
	Comments    bool          `bson:"comments" json:"comments"`
	Follow      bool          `bson:"follow" json:"follow"`
	GoodAndStar bool          `bson:"good_and_star" json:"good_and_star"`
	MailNotice  struct {
		Comments    bool `bson:"comments" json:"comments"`
		Follow      bool `bson:"follow" json:"follow"`
		GoodAndStar bool `bson:"good_and_star" json:"good_and_star"`
	} `bson:"mail_notice" json:"mail_notice"`
	UserID bson.ObjectId `bson:"user_id" json:"user_id"`
}
