package user

import (
	"github.com/minibear2333/programmer-go/api/common/result"
	"net/http"

	"github.com/minibear2333/programmer-go/api/internal/logic/user"
	"github.com/minibear2333/programmer-go/api/internal/svc"
)

func GetInterviewStatusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewGetInterviewStatusLogic(r.Context(), svcCtx)
		resp, err := l.GetInterviewStatus()
		result.HttpResult(r, w, resp, err)
	}
}
