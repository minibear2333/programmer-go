package config

type JWT struct {
	AccessSecret string `json:"AccessSecret" yaml:"AccessSecret"` // 签名
	AccessExpire int64  `json:"AccessExpire" yaml:"AccessExpire"` // 过期时间
}
