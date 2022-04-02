package interviews_tags

import (
	"github.com/minibear2333/programmer-go/api/common/result"
	"net/http"

	"github.com/minibear2333/programmer-go/api/internal/logic/interviews_tags"
	"github.com/minibear2333/programmer-go/api/internal/svc"
)

func GetAllInterviewTagsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := interviews_tags.NewGetAllInterviewTagsLogic(r.Context(), svcCtx)
		resp, err := l.GetAllInterviewTags()
		result.HttpResult(r, w, resp, err)
	}
}
