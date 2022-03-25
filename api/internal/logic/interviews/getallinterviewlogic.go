package interviews

import (
	"context"

	"github.com/minibear2333/programmer-go/api/internal/svc"
	"github.com/minibear2333/programmer-go/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllInterviewLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllInterviewLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetAllInterviewLogic {
	return GetAllInterviewLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllInterviewLogic) GetAllInterview(req types.ReqInterviews) (resp []types.Interview, err error) {
	// todo: add your logic here and delete this line

	return
}
