package interviews

import (
	"context"
	"github.com/minibear2333/programmer-go/api/global"
	"go.uber.org/zap"

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
	interviews, err := global.Mongo.InterviewsModel.FindByTagsAndSearchWord(context.TODO(), req.Tags, req.Search)
	if err != nil {
		global.LOG.Error("根据标签和关键字获取面试题列表失败", zap.Error(err))
		return nil, err
	}
	for _, interview := range *interviews {
		status := false
		for _, v := range interview.Comments {
			if v.ID.Hex() == req.UserID {
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
