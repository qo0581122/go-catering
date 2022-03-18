package logic

import (
	"context"

	"catering/coupon/rpc/internal/svc"
	proto "catering/proto/coupon"
	productBase "catering/proto/product"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserCouponLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserCouponLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserCouponLogic {
	return &GetUserCouponLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserCouponLogic) GetUserCoupon(in *proto.GetUserCouponReq) (*proto.GetUserCouponResp, error) {
	var resp proto.GetUserCouponResp
	uid := in.GetUid()
	userCoupons := l.svcCtx.CouponRepository.GetUserCoupons(uid)
	var result []*proto.UserCoupon
	for _, userCoupon := range userCoupons {
		couponId := userCoupon.GetCoupon().GetId()
		coupon := l.svcCtx.CouponRepository.GetCouponInfo(couponId)
		productId := coupon.Product.GetId()
		productResp, _ := l.svcCtx.ProductService.GetProduct(l.ctx, &productBase.GetProductReq{
			ProductId: productId,
		})
		product := productResp.GetBases()
		coupon.Product = product[0]
		userCoupon.Coupon = coupon
		result = append(result, &userCoupon)
	}
	resp.Coupons = result
	return &resp, nil
}
