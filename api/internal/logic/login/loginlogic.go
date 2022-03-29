package login

import (
	"context"
	"github.com/minibear2333/programmer-go/api/common/perr"
	"github.com/minibear2333/programmer-go/api/global"
	"github.com/minibear2333/programmer-go/api/internal/svc"
	"github.com/minibear2333/programmer-go/api/internal/types"
	"github.com/minibear2333/programmer-go/api/model"
	"github.com/minibear2333/programmer-go/api/utils"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"go.uber.org/zap"
	"strings"
	"time"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

var ErrAuthCodeError = perr.NewErrMsg("验证码错误")

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req types.ReqLogin) (resp *types.RespLogin, err error) {
	// 获取 openId
	openId := global.REDIS.HGet(context.Background(), req.ValiCode, "openid").Val()
	if len(strings.TrimSpace(openId)) == 0 {
		return nil, errors.Wrap(ErrAuthCodeError, "验证码错误")
	}

	// 登录逻辑
	user, err := global.Mongo.UserModel.FindOneByOpenId(context.TODO(), openId)
	if err != nil && err != model.ErrNotFound {
		global.LOG.Error("根据 OpenId 获取用户失败", zap.Error(err))
		return nil, errors.Wrapf(perr.NewErrCode(perr.DBError), "根据OpenId获取用户失败，OpenId: %s", openId)
	}

	if user == nil {
		// 创建新用户
		defaultUser := getDefaultUser(openId)
		if err := global.Mongo.UserModel.Insert(context.TODO(), &defaultUser); err != nil {
			global.LOG.Error("创建用户失败!", zap.Error(err))
			return nil, perr.NewErrCode(perr.ServerCommonError)
		}
		user = &defaultUser
	}

	// 颁发 Token
	j := &utils.JWT{SigningKey: []byte(global.CONFIG.Auth.AccessSecret)}
	claims := j.CreateClaims(utils.BaseClaims{
		ID:       user.ID.Hex(),
		Username: user.Name,
	})

	token, err := j.CreateToken(claims)
	if err != nil {
		global.LOG.Error("获取token失败!", zap.Error(err))
		return nil, errors.Wrapf(perr.NewErrCode(perr.TokenGenerateError), "IdentityRpc.GenerateToken userId : %s", user.ID.Hex())
	}

	return &types.RespLogin{
		ID:   user.ID.Hex(),
		Name: user.Name,
		JwtToken: types.JwtToken{
			AccessToken:  token,
			AccessExpire: claims.StandardClaims.ExpiresAt * 1000,
		},
	}, nil
}

func getDefaultUser(openId string) model.User {
	return model.User{
		OpenId:   openId,
		Avatar:   global.DefaultUserAvatar,
		Name:     global.DefaultUserName,
		Summary:  global.DefaultSummary,
		Birthday: time.Now(),
	}
}
