package perr

// OK 成功
const OK uint32 = 200

// 全局错误码
const (
	ServerCommonError  uint32 = 10001
	TokenGenerateError uint32 = 10002
	DBError            uint32 = 10003
	NotFoundError      uint32 = 10004
	NoAccessError      uint32 = 10005
	OperateFailError   uint32 = 10006
	InvalidReqError    uint32 = 10007
	InvalidParamError  uint32 = 10008
	SearchFailError    uint32 = 10009
)

// 登录模块
const (
	ErrAuthCodeError uint32 = 20001
)

//面试题模块
const (
	AuthorIdError uint32 = 30001
	InterviewGenerateError uint32 = 30002
)