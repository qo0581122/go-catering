package handler

import (
	"net/http"

	"catering/coupon/api/internal/logic"
	"catering/coupon/api/internal/svc"
	. "catering/proto/coupon"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func DrawCouponHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req DrawCouponReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewDrawCouponLogic(r.Context(), ctx)
		resp, err := l.DrawCoupon(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
