package login

import (
	"context"
	"errors"
	"github.com/minibear2333/programmer-go/api/global"
	"github.com/minibear2333/programmer-go/api/internal/svc"
	"github.com/minibear2333/programmer-go/api/internal/types"
	"github.com/minibear2333/programmer-go/api/utils"
	"go.uber.org/zap"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req types.ReqLogin) (resp *types.RespLogin, err error) {
	// 前置校验
	if len(strings.TrimSpace(req.ValiCode)) == 0 {
		return nil, errors.New("参数错误")
	}

	// 获取 openId

	// 登录逻辑

	// 颁发 Token
	j := &utils.JWT{SigningKey: []byte(global.CONFIG.Auth.AccessSecret)}
	claims := j.CreateClaims(utils.BaseClaims{
		ID:       1,
		Username: "玮子",
	})

	token, err := j.CreateToken(claims)
	if err != nil {
		global.LOG.Error("获取token失败!", zap.Error(err))
		return
	}

	return &types.RespLogin{
		ID:   "1",
		Name: "玮子",
		JwtToken: types.JwtToken{
			AccessToken:  token,
			AccessExpire: claims.StandardClaims.ExpiresAt * 1000,
		},
	}, nil
}