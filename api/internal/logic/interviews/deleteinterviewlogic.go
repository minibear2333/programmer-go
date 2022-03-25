package interviews

import (
	"context"

	"github.com/minibear2333/programmer-go/api/internal/svc"
	"github.com/minibear2333/programmer-go/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteInterviewLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteInterviewLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteInterviewLogic {
	return DeleteInterviewLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteInterviewLogic) DeleteInterview(req types.ReqInterviewId) (resp *types.CommInterviewsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
