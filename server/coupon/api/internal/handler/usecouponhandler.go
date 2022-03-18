package handler

import (
	"net/http"

	"catering/coupon/api/internal/logic"
	"catering/coupon/api/internal/svc"
	. "catering/proto/coupon"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func UseCouponHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req UseCouponReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUseCouponLogic(r.Context(), ctx)
		resp, err := l.UseCoupon(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
