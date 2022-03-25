package config

type Redis struct {
	DB       int    // redis的哪个数据库
	Addr     string // 服务器地址:端口
	Password string // 密码
}
