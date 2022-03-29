package perr

// OK 成功
const OK uint32 = 200

// 全局错误码
const (
	ServerCommonError  uint32 = 10001
	TokenGenerateError uint32 = 10002
	DBError            uint32 = 10003
)

// 登录模块
const (
	ErrAuthCodeError uint32 = 20001
)
