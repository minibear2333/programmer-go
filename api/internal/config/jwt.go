package config

type JWT struct {
	AccessSecret string // 签名
	AccessExpire int64  // 过期时间
}
