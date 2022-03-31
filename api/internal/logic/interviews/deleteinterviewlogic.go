package interviews

import (
	"context"
	"errors"
	"fmt"
	"github.com/minibear2333/programmer-go/api/global"
	"github.com/minibear2333/programmer-go/api/internal/svc"
	"github.com/minibear2333/programmer-go/api/internal/types"
	"github.com/minibear2333/programmer-go/api/model"
	"go.uber.org/zap"

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
	interview, err := global.Mongo.InterviewsModel.FindOne(l.ctx, req.ID)
	if err != nil {
		global.LOG.Error("根据ID获取面试题失败", zap.Error(err))
		return nil, err
	}
	userID := l.ctx.Value("ID")
	if interview.Author.ID.Hex() != userID{
		return nil, errors.New("无权删除该面试题目")
	}

	err = global.Mongo.InterviewsModel.Delete(l.ctx, req.ID)
	if err != nil {
		global.LOG.Error(fmt.Sprintf("删除ID为%s面试题失败", req.ID), zap.Error(err))

		if err == model.ErrNotFound {
			return &types.CommInterviewsResp{
				Ok:    false,
				Error: "不存在该面试题",
			}, nil
		}

		return &types.CommInterviewsResp{
			Ok:    false,
			Error: "删除面试题失败，请稍候重试",
		}, nil
	}

	return &types.CommInterviewsResp{
		Ok:    true,
		Error: "",
	}, nil
}
