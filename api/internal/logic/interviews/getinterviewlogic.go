package interviews

import (
	"context"
	"github.com/minibear2333/programmer-go/api/global"
	"github.com/minibear2333/programmer-go/api/internal/svc"
	"github.com/minibear2333/programmer-go/api/internal/types"
	"go.uber.org/zap"

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
	interview, err := global.Mongo.InterviewsModel.FindOne(context.TODO(), req.ID)
	if err != nil {
		global.LOG.Error("根据ID获取面试题失败", zap.Error(err))
		return nil, err
	}
	userID := l.ctx.Value("ID")

	status := false
	for _, v := range interview.Comments {
		if v.ID.Hex() == userID.(string) {
			status = true
		}
	}
	return &types.Interview_detail{
		Interview: types.Interview{
			ID: interview.ID.Hex(),
			Author: types.Author{
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
		},
		StarNum:     interview.StarNum,
		Bad:         interview.Bad,
		Content:     interview.Content,
		CreatedTime: interview.CreatedTime.Unix(),
	}, nil
}
