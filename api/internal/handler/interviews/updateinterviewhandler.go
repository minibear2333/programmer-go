package interviews

import (
	"github.com/minibear2333/programmer-go/api/common/result"
	"net/http"

	"github.com/minibear2333/programmer-go/api/internal/logic/interviews"
	"github.com/minibear2333/programmer-go/api/internal/svc"
	"github.com/minibear2333/programmer-go/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateInterviewHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ReqInterviewUpdate
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := interviews.NewUpdateInterviewLogic(r.Context(), svcCtx)
		resp, err := l.UpdateInterview(req)
		result.HttpResult(r, w, resp, err)
	}
}
