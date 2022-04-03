package result

import (
	"net/http"

	"github.com/minibear2333/programmer-go/api/common/perr"
)

type ResponseSuccess struct {
	Code uint32      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type ResponseError struct {
	Code uint32 `json:"code"`
	Msg  string `json:"msg"`
}

func Success(data interface{}) *ResponseSuccess {
	return &ResponseSuccess{http.StatusOK, perr.NewErrCode(perr.OK).GetErrMsg(), data}
}

func Error(errCode uint32, errMsg string) *ResponseError {
	return &ResponseError{errCode, errMsg}
}
