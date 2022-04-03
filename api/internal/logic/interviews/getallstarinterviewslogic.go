package interviews

import (
	"context"

	"github.com/minibear2333/programmer-go/api/internal/svc"
	"github.com/minibear2333/programmer-go/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllStarInterviewsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllStarInterviewsLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetAllStarInterviewsLogic {
	return GetAllStarInterviewsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllStarInterviewsLogic) GetAllStarInterviews(req types.ReqInterviews) (resp []types.Interview, err error) {
	// todo: add your logic here and delete this line

	return
}
