package perr

var message map[uint32]string

func init() {
	message = map[uint32]string{
		OK: "Success",
		ServerCommonError: "服务器开小差啦，稍后再来试一试吧",
		TokenGenerateError: "生成token失败",
		DBError: "数据库繁忙，请稍后再试",
		NotFoundError: "不存在相关数据",
		NoAccessError: "无权限操作",
		OperateFailError: "操作失败",
		InvalidReqError: "不合法请求",
		InvalidParamError: "不合法的参数",
		SearchFailError: "列表无数据",

		ErrAuthCodeError: "验证码错误",

		AuthorIdError: "错误的作者ID",
		InterviewGenerateError: "面试题生成失败",
	}
}

func MapErrMsg(errCode uint32) string {
	if msg, ok := message[errCode]; ok {
		return msg
	} else {
		return "服务器开小差啦，稍后再来试一试吧"
	}
}

func IsCodeErr(errCode uint32) bool {
	if _, ok := message[errCode]; ok {
		return true
	} else {
		return false
	}
}
