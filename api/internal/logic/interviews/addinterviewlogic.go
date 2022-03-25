package interviews

import (
	"context"

	"github.com/minibear2333/programmer-go/api/internal/svc"
	"github.com/minibear2333/programmer-go/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddInterviewLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddInterviewLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddInterviewLogic {
	return AddInterviewLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddInterviewLogic) AddInterview(req types.Interview_detail) (resp *types.Interview_detail, err error) {
	// todo: add your logic here and delete this line

	return
}
