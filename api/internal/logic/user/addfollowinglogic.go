package user

import (
	"context"
	"github.com/minibear2333/programmer-go/api/common/perr"
	"github.com/minibear2333/programmer-go/api/global"
	"github.com/minibear2333/programmer-go/api/internal/svc"
	"github.com/minibear2333/programmer-go/api/internal/types"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddFollowingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddFollowingLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddFollowingLogic {
	return AddFollowingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddFollowingLogic) AddFollowing(req types.ReqUserId) error {
	userID := l.ctx.Value("ID")
	err := global.Mongo.UserModel.AddUserToSetByID(l.ctx, userID.(string), "following", req.ID)
	if err != nil {
		global.LOG.Error("新增关注失败", zap.Error(err))
		return errors.Wrapf(perr.NewErrCode(perr.OperateFailError), "logic.AddFollowing fail: %s", req.ID)
	}
	err = global.Mongo.UserModel.AddUserToSetByID(l.ctx, req.ID, "followers", userID.(string))
	if err != nil {
		global.LOG.Error("新增粉丝失败", zap.Error(err))
		return errors.Wrapf(perr.NewErrCode(perr.OperateFailError), "logic.AddFollowing fail: %s", req.ID)
	}
	return nil
}
