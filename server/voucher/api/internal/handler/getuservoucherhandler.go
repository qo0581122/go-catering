package handler

import (
	"net/http"

	. "catering/proto/voucher"
	"catering/voucher/api/internal/logic"
	"catering/voucher/api/internal/svc"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func GetUserVoucherHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req GetUserVoucherReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetUserVoucherLogic(r.Context(), ctx)
		resp, err := l.GetUserVoucher(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
