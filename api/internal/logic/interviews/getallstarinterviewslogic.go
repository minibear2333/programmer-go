package interviews

import (
	"context"
	"github.com/minibear2333/programmer-go/api/common/perr"
	"github.com/minibear2333/programmer-go/api/global"
	"github.com/pkg/errors"

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
	if req.CommonPage.PageNo < global.MinPageNo {
		return nil, errors.Wrapf(perr.NewErrCode(perr.InvalidParamError), "logic.GetAllStarInterviews invalid page_no param: %d ", req.CommonPage.PageNo)
	}
	if req.CommonPage.PageSize < global.MinPageSize || req.CommonPage.PageSize > global.MaxPageSize {
		return nil, errors.Wrapf(perr.NewErrCode(perr.InvalidParamError), "logic.GetAllStarInterviews invalid page_size param: %d ", req.CommonPage.PageSize)
	}

	// TODO: ADD LOGIN
	return
}
