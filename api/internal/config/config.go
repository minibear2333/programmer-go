package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	AccessSecret string // 签名
	AccessExpire int64  // 过期时间
	Zap          Zap    `yaml:"Zap"`
	Redis        Redis  `yaml:"Redis"`
}
