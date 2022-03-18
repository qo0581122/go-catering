package handler

import (
	"net/http"

	. "catering/proto/voucher"
	"catering/voucher/api/internal/logic"
	"catering/voucher/api/internal/svc"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func GetVoucherInfoHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req GetVoucherInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetVoucherInfoLogic(r.Context(), ctx)
		resp, err := l.GetVoucherInfo(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
