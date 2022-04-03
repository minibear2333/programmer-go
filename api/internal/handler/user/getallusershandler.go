package user

import (
	"github.com/minibear2333/programmer-go/api/common/result"
	"net/http"

	"github.com/minibear2333/programmer-go/api/internal/logic/user"
	"github.com/minibear2333/programmer-go/api/internal/svc"
)

func GetAllUsersHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewGetAllUsersLogic(r.Context(), svcCtx)
		resp, err := l.GetAllUsers()
		result.HttpResult(r, w, resp, err)
	}
}
