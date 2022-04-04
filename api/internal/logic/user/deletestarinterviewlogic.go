package user

import (
	"context"

	"github.com/minibear2333/programmer-go/api/internal/svc"
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

func (l *DeleteStarInterviewLogic) DeleteStarInterview() error {
	// todo: add your logic here and delete this line

	return nil
}
