package user

import (
	"context"

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

	return nil
}
