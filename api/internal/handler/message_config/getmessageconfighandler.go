package message_config

import (
	"github.com/minibear2333/programmer-go/api/common/result"
	"net/http"

	"github.com/minibear2333/programmer-go/api/internal/logic/message_config"
	"github.com/minibear2333/programmer-go/api/internal/svc"
)

func GetMessageConfigHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := message_config.NewGetMessageConfigLogic(r.Context(), svcCtx)
		resp, err := l.GetMessageConfig()
		result.HttpResult(r, w, resp, err)
	}
}
