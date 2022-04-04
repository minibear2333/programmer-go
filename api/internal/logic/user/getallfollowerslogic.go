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

type GetAllFollowersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllFollowersLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetAllFollowersLogic {
	return GetAllFollowersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllFollowersLogic) GetAllFollowers(req types.ReqUsers) (resp []types.OtherUser, err error) {
	if req.PageNo < global.MinPageNo {
		return nil, errors.Wrapf(perr.NewErrCode(perr.InvalidParamError), "logic.GetAllFollowers invalid page_no param: %d ", req.PageNo)
	}
	if req.PageSize < global.MinPageSize || req.PageSize > global.MaxPageSize {
		return nil, errors.Wrapf(perr.NewErrCode(perr.InvalidParamError), "logic.GetAllFollowers invalid page_size param: %d ", req.PageSize)
	}
	userID := l.ctx.Value("ID")
	if req.ID != userID {
		return nil, errors.Wrapf(perr.NewErrCode(perr.NoAccessError), "logic.GetAllFollowers can't access: userId[%s] your input[%s]", userID, req.ID)
	}
	user, err := global.Mongo.UserModel.FindOne(l.ctx, req.ID)
	if err != nil {
		global.LOG.Error("获取用户失败", zap.Error(err))
		return nil, errors.Wrap(perr.NewErrCode(perr.NotFoundError), "logic.GetAllFollowers fail")
	}
	var oIDs []bson.ObjectId
	for k := range user.Followers {
		oID := user.Followers[k].ID
		oIDs = append(oIDs, oID)
	}
	users, err := global.Mongo.UserModel.FindUsersBySearchAndIds(l.ctx, req.Search, oIDs, req.PageNo, req.PageSize)
	if err != nil {
		global.LOG.Error("根据关键字和页码获取粉丝用户列表失败", zap.Error(err))
		return nil, errors.Wrapf(perr.NewErrCode(perr.SearchFailError), "logic.GetAllFollowers search users struct: %+v", req)
	}
	for k := range *users {
		u := (*users)[k]
		tmp := types.OtherUser{
			ID:      u.ID.Hex(),
			Avatar:  u.Avatar,
			Name:    u.Name,
			Summary: u.Summary,
		}
		resp = append(resp, tmp)
	}
	return
}
