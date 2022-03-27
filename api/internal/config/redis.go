package config

type Redis struct {
	DB       int    `json:"DB" yaml:"DB"`             // redis的哪个数据库
	Addr     string `json:"Addr" yaml:"Addr"`         // 服务器地址:端口
	Password string `json:"Password" yaml:"Password"` // 密码
}
