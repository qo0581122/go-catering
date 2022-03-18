package logic

import (
	"context"

	"catering/coupon/api/internal/svc"
	. "catering/proto/coupon"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserCouponLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserCouponLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetUserCouponLogic {
	return GetUserCouponLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserCouponLogic) GetUserCoupon(req GetUserCouponReq) (*GetUserCouponResp, error) {
	return l.svcCtx.ConponRpc.GetUserCoupon(l.ctx, &req)
}
