package interviews

import (
	"github.com/minibear2333/programmer-go/api/common/result"
	"net/http"

	"github.com/minibear2333/programmer-go/api/internal/logic/interviews"
	"github.com/minibear2333/programmer-go/api/internal/svc"
	"github.com/minibear2333/programmer-go/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteInterviewHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ReqInterviewId
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := interviews.NewDeleteInterviewLogic(r.Context(), svcCtx)
		resp, err := l.DeleteInterview(req)
		result.HttpResult(r, w, resp, err)
	}
}
