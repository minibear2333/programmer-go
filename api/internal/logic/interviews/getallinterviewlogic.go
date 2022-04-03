package interviews

import (
	"context"
	"github.com/minibear2333/programmer-go/api/common/perr"
	"github.com/pkg/errors"

	"github.com/minibear2333/programmer-go/api/global"
	"github.com/minibear2333/programmer-go/api/internal/svc"
	"github.com/minibear2333/programmer-go/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"go.uber.org/zap"
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
	if req.CommonPage.PageNo < global.MinPageNo {
		return nil, errors.Wrapf(perr.NewErrCode(perr.InvalidParamError), "logic.getAllInterview invalid page_no param: %d ", req.CommonPage.PageNo)
	}
	if req.CommonPage.PageSize < global.MinPageSize || req.CommonPage.PageSize > global.MaxPageSize {
		return nil, errors.Wrapf(perr.NewErrCode(perr.InvalidParamError), "logic.getAllInterview invalid page_size param: %d ", req.CommonPage.PageSize)
	}

	interviews, err := global.Mongo.InterviewsModel.FindByTagsAndSearchWord(context.TODO(), req.Tags, req.Search, req.CommonPage)
	if err != nil {
		global.LOG.Error("根据标签和关键字获取面试题列表失败", zap.Error(err))
		return nil, errors.Wrapf(perr.NewErrCode(perr.SearchFailError), "logic.getAllInterview search interview struct: %+v", req)
	}
	userID := l.ctx.Value("ID")

	resp = []types.Interview{}
	for _, interview := range *interviews {
		status := false
		for _, v := range interview.Comments {
			if v.ID.Hex() == userID.(string) {
				status = true
			}
		}
		tmp := types.Interview{
			ID:          interview.ID.Hex(),
			Author:      types.Author{
				ID:   interview.Author.ID.Hex(),
				Name: interview.Author.Name,
			},
			ClickNum:    interview.ClickNum,
			Good:        interview.Good,
			HardStatus:  interview.HardStatus,
			HotNum:      interview.HotNum,
			Summary:     interview.Summary,
			Tags:        interview.Tags,
			Title:       interview.Title,
			UpdatedTime: interview.UpdatedTime.Unix(),
			Status:      status,
		}
		resp = append(resp, tmp)
	}

	return resp, nil
}
