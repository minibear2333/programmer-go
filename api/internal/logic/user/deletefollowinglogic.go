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

type DeleteFollowingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteFollowingLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteFollowingLogic {
	return DeleteFollowingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteFollowingLogic) DeleteFollowing(req types.ReqUserId) error {
	userID := l.ctx.Value("ID")
	err := global.Mongo.UserModel.DeleteUserToSetByID(l.ctx, userID.(string), "following", req.ID)
	if err != nil {
		global.LOG.Error("取消关注失败", zap.Error(err))
		return errors.Wrapf(perr.NewErrCode(perr.OperateFailError), "logic.DeleteFollowing fail: %s", req.ID)
	}
	err = global.Mongo.UserModel.DeleteUserToSetByID(l.ctx,req.ID, "followers",  userID.(string))
	if err != nil {
		global.LOG.Error("移除粉丝记录失败", zap.Error(err))
		return errors.Wrapf(perr.NewErrCode(perr.OperateFailError), "logic.DeleteFollowing fail: %s", req.ID)
	}
	return nil
}
