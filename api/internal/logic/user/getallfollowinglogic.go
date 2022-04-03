package user

import (
	"context"

	"github.com/minibear2333/programmer-go/api/internal/svc"
	"github.com/minibear2333/programmer-go/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllFollowingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllFollowingLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetAllFollowingLogic {
	return GetAllFollowingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllFollowingLogic) GetAllFollowing(req types.ReqUsers) (resp []types.OtherUser, err error) {
	// todo: add your logic here and delete this line

	return
}
