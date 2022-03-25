package interviews

import (
	"context"

	"github.com/minibear2333/programmer-go/api/internal/svc"
	"github.com/minibear2333/programmer-go/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetInterviewLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetInterviewLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetInterviewLogic {
	return GetInterviewLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetInterviewLogic) GetInterview(req types.ReqInterviewId) (resp *types.Interview_detail, err error) {
	// todo: add your logic here and delete this line

	return
}
