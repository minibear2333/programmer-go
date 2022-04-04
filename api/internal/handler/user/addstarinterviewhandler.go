package user

import (
	"github.com/minibear2333/programmer-go/api/common/result"
	"net/http"

	"github.com/minibear2333/programmer-go/api/internal/logic/user"
	"github.com/minibear2333/programmer-go/api/internal/svc"
)

func AddStarInterviewHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewAddStarInterviewLogic(r.Context(), svcCtx)
		err := l.AddStarInterview()
		result.HttpResult(r, w, nil, err)
	}
}
