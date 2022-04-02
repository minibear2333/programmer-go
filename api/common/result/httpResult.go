package result

import (
	"github.com/minibear2333/programmer-go/api/common/perr"
	"net/http"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// HttpResult http返回
func HttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	if err == nil {
		// 成功返回
		r := Success(resp)
		httpx.WriteJson(w, http.StatusOK, r)
	} else {
		// 错误返回
		errCode := perr.ServerCommonError
		errMsg := "服务器开小差啦，稍后再来试一试"

		causeErr := errors.Cause(err)                // err类型
		if e, ok := causeErr.(*perr.CodeError); ok { // 自定义错误类型
			//自定义 CodeError
			errCode = e.GetErrCode()
			errMsg = e.GetErrMsg()
		}
		logx.WithContext(r.Context()).Errorf("【API-ERR】 : %+v ", err)
		httpx.WriteJson(w, http.StatusBadRequest, Error(errCode, errMsg))
	}
}
