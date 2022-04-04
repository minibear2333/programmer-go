package user

import (
	"context"

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
	// todo: add your logic here and delete this line

	return nil
}
