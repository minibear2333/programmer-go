package message_config

import (
	"context"
	"github.com/globalsign/mgo/bson"
	"github.com/minibear2333/programmer-go/api/common/perr"
	"github.com/minibear2333/programmer-go/api/global"
	"github.com/minibear2333/programmer-go/api/model"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/minibear2333/programmer-go/api/internal/svc"
	"github.com/minibear2333/programmer-go/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMessageConfigLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateMessageConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateMessageConfigLogic {
	return UpdateMessageConfigLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateMessageConfigLogic) UpdateMessageConfig(req types.MessageConfig) (resp *types.MessageConfig, err error) {
	userID := l.ctx.Value("ID")
	m := model.MessageConfig{
		Comments:    req.Comments,
		Follow:      req.Follow,
		GoodAndStar: req.GoodAndStar,
		MailNotice: struct {
			Comments    bool `bson:"comments" json:"comments"`
			Follow      bool `bson:"follow" json:"follow"`
			GoodAndStar bool `bson:"good_and_star" json:"good_and_star"`
		}{
			Comments:    req.MailNotice.Comments,
			Follow:      req.MailNotice.Follow,
			GoodAndStar: req.MailNotice.GoodAndStar,
		},
		UserID: bson.ObjectIdHex(userID.(string)),
	}

	messageConfig, err := global.Mongo.MessageConfigModel.FindOneByUserID(l.ctx, userID.(string))
	if messageConfig == nil && err == model.ErrNotFound {
		err = global.Mongo.MessageConfigModel.Insert(l.ctx, &m)
		if err != nil {
			global.LOG.Error("创建消息中心配置失败", zap.Error(err))
			return nil, perr.NewErrCode(perr.MessageConfigGenerateError)
		}
		return &req, nil
	}
	if err != nil {
		global.LOG.Error("通过用户id获取消息中心配置失败", zap.Error(err))
		return nil, errors.Wrapf(perr.NewErrCode(perr.NotFoundError), "logic.UpdateMessageConfig.FindOneByUserID not found: %s", userID)
	}

	messageConfig.Comments = req.Comments
	messageConfig.Follow = req.Follow
	messageConfig.GoodAndStar = req.GoodAndStar
	messageConfig.MailNotice = struct {
		Comments    bool `bson:"comments" json:"comments"`
		Follow      bool `bson:"follow" json:"follow"`
		GoodAndStar bool `bson:"good_and_star" json:"good_and_star"`
	}{
		Comments:    req.MailNotice.Comments,
		Follow:      req.MailNotice.Follow,
		GoodAndStar: req.MailNotice.GoodAndStar,
	}

	err = global.Mongo.MessageConfigModel.UpdateByUserID(l.ctx, messageConfig)
	if err != nil {
		global.LOG.Error("更新消息中心配置失败:", zap.Error(err))
		return nil, errors.Wrapf(perr.NewErrCode(perr.OperateFailError), "logic.UpdateMessageConfig operate fail: %s", userID)
	}
	return &req, nil
}
