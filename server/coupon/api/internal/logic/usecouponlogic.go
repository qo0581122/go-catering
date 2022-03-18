package logic

import (
	"context"

	"catering/coupon/api/internal/svc"
	. "catering/proto/coupon"

	"github.com/tal-tech/go-zero/core/logx"
)

type UseCouponLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUseCouponLogic(ctx context.Context, svcCtx *svc.ServiceContext) UseCouponLogic {
	return UseCouponLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UseCouponLogic) UseCoupon(req UseCouponReq) (*UseCouponResp, error) {
	return l.svcCtx.ConponRpc.UseCoupon(l.ctx, &req)
}
