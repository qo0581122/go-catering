package handler

import (
	"net/http"

	"catering/advertisement/api/internal/logic"
	"catering/advertisement/api/internal/svc"
	"catering/advertisement/api/internal/types"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func GetHotActivitieHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetAdHotActivitieReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetHotActivitieLogic(r.Context(), ctx)
		resp, err := l.GetHotActivitie(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
