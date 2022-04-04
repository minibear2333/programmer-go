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

type DeleteStarInterviewLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteStarInterviewLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteStarInterviewLogic {
	return DeleteStarInterviewLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteStarInterviewLogic) DeleteStarInterview(req types.ReqStarInterviewsID) error {
	userID := l.ctx.Value("ID")
	err := global.Mongo.UserModel.DeleteUserToSetByID(l.ctx, userID.(string), "star_interviews", req.ID)
	if err != nil {
		global.LOG.Error("取消收藏面试题失败", zap.Error(err))
		return errors.Wrapf(perr.NewErrCode(perr.OperateFailError), "logic.DeleteStarInterview fail: %s", req.ID)
	}
	return nil
}
