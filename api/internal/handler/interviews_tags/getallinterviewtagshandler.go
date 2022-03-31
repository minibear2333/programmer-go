package interviews_tags

import (
	"net/http"

	"github.com/minibear2333/programmer-go/api/internal/logic/interviews_tags"
	"github.com/minibear2333/programmer-go/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetAllInterviewTagsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := interviews_tags.NewGetAllInterviewTagsLogic(r.Context(), svcCtx)
		resp, err := l.GetAllInterviewTags()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
