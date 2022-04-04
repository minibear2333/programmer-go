package user

import (
	"context"
	"github.com/globalsign/mgo/bson"
	"github.com/minibear2333/programmer-go/api/common/perr"
	"github.com/minibear2333/programmer-go/api/global"
	"github.com/pkg/errors"
	"go.uber.org/zap"

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

func (l *GetAllStarInterviewsLogic) GetAllStarInterviews(req types.ReqStarInterviews) (resp []types.RespStarInterviews, err error) {
	if req.PageNo < global.MinPageNo {
		return nil, errors.Wrapf(perr.NewErrCode(perr.InvalidParamError), "logic.GetAllStarInterviews invalid page_no param: %d ", req.PageNo)
	}
	if req.PageSize < global.MinPageSize || req.PageSize > global.MaxPageSize {
		return nil, errors.Wrapf(perr.NewErrCode(perr.InvalidParamError), "logic.GetAllStarInterviews invalid page_size param: %d ", req.PageSize)
	}
	user, err := global.Mongo.UserModel.FindOne(l.ctx, req.ID)
	if err != nil {
		global.LOG.Error("获取用户失败", zap.Error(err))
		return nil, errors.Wrap(perr.NewErrCode(perr.NotFoundError), "logic.GetAllStarInterviews fail")
	}
	var oIDs []bson.ObjectId
	for k := range user.StarInterviews {
		oID := user.StarInterviews[k].ID
		oIDs = append(oIDs, oID)
	}
	interviews, err := global.Mongo.InterviewsModel.FindByTagsAndSearchWord(l.ctx, oIDs, req.Tags, req.Search, types.CommonPage{
		PageNo:   req.PageNo,
		PageSize: req.PageSize,
	})

	for k := range *interviews {
		interview := (*interviews)[k]
		var tmp = types.RespStarInterviews{
			ID:          interview.ID.Hex(),
			Title:       interview.Title,
			UpdatedTime: interview.UpdatedTime.Unix(),
		}
		resp = append(resp, tmp)
	}
	return
}
