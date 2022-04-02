package interviews

import (
	"context"
	"fmt"
	"github.com/minibear2333/programmer-go/api/common/perr"
	"github.com/minibear2333/programmer-go/api/global"
	"github.com/minibear2333/programmer-go/api/internal/svc"
	"github.com/minibear2333/programmer-go/api/internal/types"
	"github.com/pkg/errors"
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

func (l *DeleteInterviewLogic) DeleteInterview(req types.ReqInterviewId) (res bool,  err error) {
	interview, err := global.Mongo.InterviewsModel.FindOne(l.ctx, req.ID)
	if err != nil {
		global.LOG.Error("根据ID获取面试题失败", zap.Error(err))
		return false, errors.Wrapf(perr.NewErrCode(perr.NotFoundError), "logic.deleteIntervirews not found: %s", req.ID)
	}
	userID := l.ctx.Value("ID")
	if interview.Author.ID.Hex() != userID{
		return false, errors.Wrapf(perr.NewErrCode(perr.NotFoundError), "logic.deleteIntervirews not access: userId[%s] interviewId[%s]", userID, req.ID)
	}

	err = global.Mongo.InterviewsModel.Delete(l.ctx, req.ID)
	if err != nil {
		global.LOG.Error(fmt.Sprintf("删除面试题失败:%s", req.ID), zap.Error(err))
		return false, errors.Wrapf(perr.NewErrCode(perr.OperateFailError), "logic.deleteIntervirews operate fail: interviewId[%s]", userID, req.ID)
	}

	return true, nil
}
