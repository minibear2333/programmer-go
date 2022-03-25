package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/minibear2333/programmer-go/api/global"
	"time"
)

type CustomClaims struct {
	BaseClaims
	jwt.StandardClaims
}

type BaseClaims struct {
	ID       uint
	Username string
}

type JWT struct {
	SigningKey []byte
}

func (j *JWT) CreateClaims(baseClaims BaseClaims) CustomClaims {
	claims := CustomClaims{
		BaseClaims: baseClaims,
		// TODO 缓冲令牌
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                       // 签名生效时间
			ExpiresAt: time.Now().Unix() + global.CONFIG.AccessExpire, // 过期时间 7天  配置文件
		},
	}
	return claims
}

func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}
