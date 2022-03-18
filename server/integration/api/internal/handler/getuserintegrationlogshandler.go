package handler

import (
	"net/http"

	"catering/integration/api/internal/logic"
	"catering/integration/api/internal/svc"
	. "catering/proto/integration"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func GetUserIntegrationLogsHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req GetUserIntegrationChangeLogsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetUserIntegrationLogsLogic(r.Context(), ctx)
		resp, err := l.GetUserIntegrationLogs(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
