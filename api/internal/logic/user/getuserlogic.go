package user

import (
	"context"
	"github.com/globalsign/mgo/bson"
	"github.com/minibear2333/programmer-go/api/common/perr"
	"github.com/minibear2333/programmer-go/api/global"
	"github.com/minibear2333/programmer-go/api/internal/svc"
	"github.com/minibear2333/programmer-go/api/internal/types"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetUserLogic {
	return GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req types.ReqUserId) (resp *types.RespUser, err error) {
	if !bson.IsObjectIdHex(req.ID) {
		global.LOG.Error("id识别错误")
		return nil, errors.Wrapf(perr.NewErrCode(perr.InvalidParamError), "logic.GetUser invalid param: %s", req.ID)
	}
	user, err := global.Mongo.UserModel.FindOne(l.ctx, req.ID)
	if err != nil {
		global.LOG.Error("获取用户失败", zap.Error(err))
		return nil, errors.Wrapf(perr.NewErrCode(perr.NotFoundError), "logic.GetUser fail")
	}
	userID := l.ctx.Value("ID")
	if user.ID.Hex() != userID {
		return nil, errors.Wrapf(perr.NewErrCode(perr.NoAccessError), "logic.GetUser can't access: userId[%s] your input[%s]", userID, req.ID)
	}

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
