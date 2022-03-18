package handler

import (
	"net/http"

	"catering/area/api/internal/logic"
	"catering/area/api/internal/svc"
	"catering/area/api/internal/types"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func cityHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetCitysReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCityLogic(r.Context(), ctx)
		resp, err := l.City(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
