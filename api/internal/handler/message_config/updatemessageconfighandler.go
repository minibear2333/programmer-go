package message_config

import (
	"github.com/minibear2333/programmer-go/api/common/perr"
	"github.com/minibear2333/programmer-go/api/common/result"
	"github.com/pkg/errors"
	"net/http"

	"github.com/minibear2333/programmer-go/api/internal/logic/message_config"
	"github.com/minibear2333/programmer-go/api/internal/svc"
	"github.com/minibear2333/programmer-go/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateMessageConfigHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MessageConfig
		if err := httpx.Parse(r, &req); err != nil {
			err = errors.Wrap(perr.NewErrCode(perr.InvalidParamError), err.Error())
			result.HttpResult(r, w, nil, err)
			return
		}

		l := message_config.NewUpdateMessageConfigLogic(r.Context(), svcCtx)
		resp, err := l.UpdateMessageConfig(req)
		result.HttpResult(r, w, resp, err)
	}
}
