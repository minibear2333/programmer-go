package user

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

type UpdateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateUserLogic {
	return UpdateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserLogic) UpdateUser(req types.ReqUser) (resp *types.RespUser, err error) {
	if !bson.IsObjectIdHex(req.ID) {
		global.LOG.Error("id识别错误")
		return nil, errors.Wrapf(perr.NewErrCode(perr.InvalidParamError), "logic.UpdateUser invalid param: %s", req.ID)
	}
	user, err := global.Mongo.UserModel.FindOne(l.ctx, req.ID)
	if err != nil {
		global.LOG.Error("获取用户失败", zap.Error(err))
		return nil, errors.Wrapf(perr.NewErrCode(perr.NotFoundError), "logic.UpdateUser fail")
	}
	userID := l.ctx.Value("ID")
	if user.ID.Hex() != userID {
		return nil, errors.Wrapf(perr.NewErrCode(perr.NoAccessError), "logic.UpdateUser can't access: userId[%s] your input[%s]", userID, req.ID)
	}
	if req.Summary != nil {
		user.Summary = *req.Summary
	}
	if req.Name != nil {
		user.Name = *req.Name
	}
	if req.Avatar != nil {
		user.Avatar = *req.Avatar
	}
	if req.Blog != nil {
		user.Blog = *req.Blog
	}
	if req.City != nil {
		user.City = *req.City
	}
	if req.Email != nil {
		user.Email = *req.Email
	}
	if req.Phone != nil {
		user.Phone = *req.Phone
	}
	if req.Skills != nil {
		user.Skills = req.Skills
	}
	if req.RealName != nil {
		user.RealName = *req.RealName
	}
	if req.Birthday != nil {
		user.Birthday = time.Unix(*req.Birthday,0)
	}
	err = global.Mongo.UserModel.Update(l.ctx,user)

	return &types.RespUser{
		ID:        user.ID.Hex(),
		Avatar:    user.Avatar,
		Birthday:  user.Birthday.Unix(),
		Blog:      user.Blog,
		City:      user.City,
		Email:     user.Email,
		Followers: int64(len(user.Followers)),
		Following: int64(len(user.Following)),
		Name:      user.Name,
		Phone:     user.Phone,
		RealName:  user.RealName,
		Skills:    user.Skills,
		Summary:   user.Summary,
	}, nil
}
