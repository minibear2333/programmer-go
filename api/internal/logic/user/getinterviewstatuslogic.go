package user

import (
	"context"
	"github.com/minibear2333/programmer-go/api/common/perr"
	"github.com/minibear2333/programmer-go/api/global"
	"github.com/minibear2333/programmer-go/api/model"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/minibear2333/programmer-go/api/internal/svc"
	"github.com/minibear2333/programmer-go/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetInterviewStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetInterviewStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetInterviewStatusLogic {
	return GetInterviewStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetInterviewStatusLogic) GetInterviewStatus(req types.ReqUserId) (resp []types.InterviewHardStatus, err error) {
	user, err := global.Mongo.UserModel.FindOne(l.ctx, req.ID)
	if err != nil {
		global.LOG.Error("获取用户失败", zap.Error(err))
		return nil, errors.Wrap(perr.NewErrCode(perr.NotFoundError), "logic.GetInterviewStatus fail")
	}
	var interviewsStatus []model.CountResult
	err = global.Mongo.InterviewsModel.CountHardStatus(l.ctx, &interviewsStatus)
	if err != nil {
		global.LOG.Error("获取总体面试题难度数量失败", zap.Error(err))
		return nil, errors.Wrap(perr.NewErrCode(perr.NotFoundError), "logic.GetInterviewStatus fail")
	}

	for k := range interviewsStatus {
		difficulty := interviewsStatus[k].ID
		tmp := types.InterviewHardStatus{
			Difficulty: difficulty,
			Count:      user.InterviewsStatus[difficulty],
			Total:      interviewsStatus[k].Count,
		}
		resp = append(resp, tmp)
	}
	return
}
