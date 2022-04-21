package message_config

import (
	"context"
	"github.com/minibear2333/programmer-go/api/common/perr"
	"github.com/minibear2333/programmer-go/api/global"
	"github.com/minibear2333/programmer-go/api/model"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/minibear2333/programmer-go/api/internal/svc"
	"github.com/minibear2333/programmer-go/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMessageConfigLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMessageConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetMessageConfigLogic {
	return GetMessageConfigLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMessageConfigLogic) GetMessageConfig() (resp *types.MessageConfig, err error) {
	resp = &types.MessageConfig{
		Comments:    true,
		Follow:      true,
		GoodAndStar: true,
		MailNotice: types.MailNotice{
			Comments:    true,
			Follow:      true,
			GoodAndStar: true,
		},
	}
	userID := l.ctx.Value("ID")
	messageConfig, err := global.Mongo.MessageConfigModel.FindOneByUserID(l.ctx, userID.(string))
	if messageConfig == nil && err == model.ErrNotFound {
		return resp,nil
	}
	if err != nil {
		global.LOG.Error("通过用户id获取消息中心配置失败", zap.Error(err))
		return nil, errors.Wrapf(perr.NewErrCode(perr.NotFoundError), "logic.GetMessageConfig not found: %s", userID)
	}
	return &types.MessageConfig{
		Comments:    messageConfig.Comments,
		Follow:      messageConfig.Follow,
		GoodAndStar: messageConfig.GoodAndStar,
		MailNotice: types.MailNotice{
			Comments:    messageConfig.MailNotice.Comments,
			Follow:      messageConfig.MailNotice.Follow,
			GoodAndStar: messageConfig.MailNotice.GoodAndStar,
		},
	}, nil
}
