package handler

import (
	"net/http"

	. "catering/proto/voucher"
	"catering/voucher/api/internal/logic"
	"catering/voucher/api/internal/svc"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func UseVoucherHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req UseVoucherReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUseVoucherLogic(r.Context(), ctx)
		resp, err := l.UseVoucher(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
