package interviews

import (
	"context"
	"github.com/globalsign/mgo/bson"
	"github.com/minibear2333/programmer-go/api/common/perr"
	"github.com/minibear2333/programmer-go/api/global"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"time"

	"github.com/minibear2333/programmer-go/api/internal/svc"
	"github.com/minibear2333/programmer-go/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateInterviewLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateInterviewLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateInterviewLogic {
	return UpdateInterviewLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateInterview  这个接口应该抛弃，不应该全量更新
func (l *UpdateInterviewLogic) UpdateInterview(req types.ReqInterviewUpdate) (resp *types.Interview_detail, err error) {
	// 校验标签输入
	is_ok,err := global.Mongo.InterviewsTagsModel.CheckTag(l.ctx,req.Tags)
	if err != nil {
		global.LOG.Error("创建面试题失败", zap.Error(err))
		return nil, perr.NewErrCode(perr.InvalidInterviewTags)
	}
	if !is_ok{
		return nil, perr.NewErrCode(perr.InvalidInterviewTags)
	}
	if !bson.IsObjectIdHex(req.ID) {
		global.LOG.Error("面试题目id识别错误")
		return nil, errors.Wrapf(perr.NewErrCode(perr.InvalidParamError), "logic.updateInterview invalid param: %s", req.ID)
	}
	interview, err := global.Mongo.InterviewsModel.FindOne(l.ctx, req.ID)
	if err != nil {
		global.LOG.Error("根据ID获取面试题失败", zap.Error(err))
		return nil, errors.Wrapf(perr.NewErrCode(perr.NotFoundError), "logic.updateInterview not found: %s", req.ID)
	}
	userID := l.ctx.Value("ID")
	if interview.Author.ID.Hex() != userID{
		return nil, errors.Wrapf(perr.NewErrCode(perr.NoAccessError), "logic.updateInterview not access: userId[%s] interviewId[%s]", userID, req.ID)
	}

	interview.UpdatedTime = time.Now()
	if req.HardStatus!=nil{
		interview.HardStatus = *req.HardStatus
	}
	if req.Summary!=nil{
		interview.Summary = *req.Summary
	}
	if req.Tags!=nil{
		interview.Tags = req.Tags
	}
	if req.Title!=nil{
		interview.Title = *req.Title
	}
	if req.Content !=nil{
		interview.Title = *req.Content
	}

	err = global.Mongo.InterviewsModel.Update(l.ctx, interview)
	if err != nil {
		global.LOG.Error("更新面试题失败:", zap.Error(err))
		return nil, errors.Wrapf(perr.NewErrCode(perr.OperateFailError), "logic.updateInterview operate fail: %s", req.ID)
	}

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
