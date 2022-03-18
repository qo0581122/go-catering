package logic

import (
	"context"

	"catering/coupon/api/internal/svc"
	. "catering/proto/coupon"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetCouponInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCouponInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetCouponInfoLogic {
	return GetCouponInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCouponInfoLogic) GetCouponInfo(req GetCouponInfoReq) (*GetCouponInfoResp, error) {
	return l.svcCtx.ConponRpc.GetCouponInfo(l.ctx, &req)
}
