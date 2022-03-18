package logic

import (
	"context"

	"catering/area/pkg/errors"
	"catering/coupon/rpc/internal/svc"
	proto "catering/proto/coupon"
	productBase "catering/proto/product"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetCouponInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCouponInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCouponInfoLogic {
	return &GetCouponInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCouponInfoLogic) GetCouponInfo(in *proto.GetCouponInfoReq) (*proto.GetCouponInfoResp, error) {
	resp := &proto.GetCouponInfoResp{
		Coupon: &proto.Coupon{},
	}
	couponId := in.GetCouponId()
	if couponId == 0 {
		return nil, errors.InternalServerError("coupon id not zero")
	}
	coupon := l.svcCtx.CouponRepository.GetCouponInfo(couponId)
	productId := coupon.Product.GetId()
	productResp, _ := l.svcCtx.ProductService.GetProduct(l.ctx, &productBase.GetProductReq{
		ProductId: productId,
	})
	product := productResp.GetBases()
	coupon.Product = product[0]
	resp.Coupon = coupon
	return resp, nil
}
