package logic

import (
	"context"

	"catering/coupon/api/internal/svc"
	. "catering/proto/coupon"

	"github.com/tal-tech/go-zero/core/logx"
)

type DrawCouponLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDrawCouponLogic(ctx context.Context, svcCtx *svc.ServiceContext) DrawCouponLogic {
	return DrawCouponLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DrawCouponLogic) DrawCoupon(req DrawCouponReq) (*DrawCouponResp, error) {
	return l.svcCtx.ConponRpc.DrawCoupon(l.ctx, &req)
}
