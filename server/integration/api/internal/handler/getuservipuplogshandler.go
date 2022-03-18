package handler

import (
	"net/http"

	"catering/integration/api/internal/logic"
	"catering/integration/api/internal/svc"
	. "catering/proto/integration"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func GetUserVipUpLogsHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req GetUserVipLevelUpLogsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetUserVipUpLogsLogic(r.Context(), ctx)
		resp, err := l.GetUserVipUpLogs(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
