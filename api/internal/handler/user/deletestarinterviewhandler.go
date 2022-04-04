package user

import (
	"github.com/minibear2333/programmer-go/api/common/result"
	"net/http"

	"github.com/minibear2333/programmer-go/api/internal/logic/user"
	"github.com/minibear2333/programmer-go/api/internal/svc"
)

func DeleteStarInterviewHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewDeleteStarInterviewLogic(r.Context(), svcCtx)
		err := l.DeleteStarInterview()
		result.HttpResult(r, w, nil, err)
	}
}
