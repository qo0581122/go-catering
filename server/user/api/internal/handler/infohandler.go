package handler

import (
	"net/http"

	"catering/user/api/internal/logic"
	"catering/user/api/internal/svc"
	"catering/user/api/internal/types"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func infoHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewInfoLogic(r.Context(), ctx)
		resp, err := l.Info(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
