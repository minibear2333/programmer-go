package interviews_tags

import (
	"context"
	"github.com/minibear2333/programmer-go/api/global"
	"go.uber.org/zap"

	"github.com/minibear2333/programmer-go/api/internal/svc"
	"github.com/minibear2333/programmer-go/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllInterviewTagsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllInterviewTagsLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetAllInterviewTagsLogic {
	return GetAllInterviewTagsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllInterviewTagsLogic) GetAllInterviewTags() (resp []types.RespInterviewsTags, err error) {
	tags, err := global.Mongo.InterviewsTagsModel.FindAll(l.ctx)
	if err != nil {
		global.LOG.Error("获取面试题标签失败", zap.Error(err))
		return nil, err
	}

	for _, tag := range *tags {
		tmp := types.RespInterviewsTags{
			ID:      tag.ID.Hex(),
			Name:    tag.Name,
			SubTags: tag.SubTags,
		}
		resp = append(resp, tmp)
	}
	return resp, nil
}
