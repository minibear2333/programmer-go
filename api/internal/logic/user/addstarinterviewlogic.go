package user

import (
	"context"
	"github.com/minibear2333/programmer-go/api/common/perr"
	"github.com/minibear2333/programmer-go/api/global"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/minibear2333/programmer-go/api/internal/svc"
	"github.com/minibear2333/programmer-go/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddStarInterviewLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddStarInterviewLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddStarInterviewLogic {
	return AddStarInterviewLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddStarInterviewLogic) AddStarInterview(req types.ReqStarInterviewsID) error {
	userID := l.ctx.Value("ID")
	err := global.Mongo.UserModel.AddUserToSetByID(l.ctx, userID.(string), "star_interviews", req.ID)
	if err != nil {
		global.LOG.Error("收藏面试题失败", zap.Error(err))
		return errors.Wrapf(perr.NewErrCode(perr.OperateFailError), "logic.AddStarInterview fail: %s", req.ID)
	}
	return nil
}
