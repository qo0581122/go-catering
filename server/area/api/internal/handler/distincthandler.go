package handler

import (
	"net/http"

	"catering/area/api/internal/logic"
	"catering/area/api/internal/svc"
	"catering/area/api/internal/types"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func distinctHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetDistinctsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewDistinctLogic(r.Context(), ctx)
		resp, err := l.Distinct(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
