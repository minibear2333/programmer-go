package interviews

import (
	"context"

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

func (l *AddStarInterviewLogic) AddStarInterview(req types.ReqInterviewId) error {
	// todo: add your logic here and delete this line

	return nil
}
