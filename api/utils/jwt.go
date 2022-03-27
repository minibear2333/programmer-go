package utils

import (
	"github.com/golang-jwt/jwt"
	"github.com/minibear2333/programmer-go/api/global"
	"time"
)

type CustomClaims struct {
	BaseClaims
	jwt.StandardClaims
}

type BaseClaims struct {
	ID       string
	Username string
}

type JWT struct {
	SigningKey []byte
}

func (j *JWT) CreateClaims(baseClaims BaseClaims) CustomClaims {
	authCfg := global.CONFIG.Auth
	claims := CustomClaims{
		BaseClaims: baseClaims,
		// TODO 缓冲令牌
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                 // 签名生效时间
			ExpiresAt: time.Now().Unix() + authCfg.AccessExpire, // 过期时间 7天  配置文件
		},
	}
	return claims
}

func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}
