package perr

var message map[uint32]string

func init() {
	message = make(map[uint32]string)
	message[OK] = "Success"
	message[ServerCommonError] = "服务器开小差啦，稍后再来试一试吧"
	message[TokenGenerateError] = "生成token失败"
	message[ErrAuthCodeError] = "验证码错误"
	message[DBError] = "数据库繁忙，请稍后再试"
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
