package initialize

import (
	"context"
	"github.com/minibear2333/programmer-go/api/global"
	"github.com/minibear2333/programmer-go/api/model"
	"go.uber.org/zap"
)

func Mongo() {
	mongoCfg := global.CONFIG.Mongo
	global.Mongo = &model.MongoClient{
		CommentsModel:       model.NewCommentsModel(mongoCfg.Addr, "comments"),
		InterviewsModel:     model.NewInterviewsModel(mongoCfg.Addr, "interviews"),
		InterviewsTagsModel: model.NewInterviewsTagsModel(mongoCfg.Addr, "interviews_tags"),
		UserModel:           model.NewUserModel(mongoCfg.Addr, "user"),
		MessageCenterModel:  model.NewMessageCenterModel(mongoCfg.Addr, "message_center"),
		MessageConfigModel:  model.NewMessageConfigModel(mongoCfg.Addr, "message_config"),
	}
	_, err := global.Mongo.UserModel.FindOne(context.TODO(), "abc")
	if err == model.ErrInvalidObjectId {
		global.LOG.Info("mongo connect success!")
	} else {
		global.LOG.Info("mongo connect error :", zap.Error(err))
	}
}
