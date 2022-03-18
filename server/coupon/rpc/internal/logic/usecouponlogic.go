package logic

import (
	"context"

	"catering/area/pkg/errors"
	"catering/coupon/rpc/internal/svc"
	. "catering/coupon/rpc/repository"
	proto "catering/proto/coupon"

	"github.com/tal-tech/go-zero/core/logx"
)

type UseCouponLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUseCouponLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UseCouponLogic {
	return &UseCouponLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UseCouponLogic) UseCoupon(in *proto.UseCouponReq) (*proto.UseCouponResp, error) {
	resp := &proto.UseCouponResp{}
	var (
		userCouponId = in.GetUserCouponId()
		uid          = in.GetUid()
		orderId      = in.GetOrderId()
	)
	if userCouponId == 0 || uid == 0 || orderId == 0 {
		return resp, errors.InternalServerError("id no zero")
	}
	userCoupon := l.svcCtx.CouponRepository.GetUserCouponsById(userCouponId)
	if err := CheckUseStatus(userCoupon.UseStatus); err != nil {
		return resp, err
	}
	if err := CheckUseTime(userCoupon.GetUseBeginTime(), userCoupon.GetUseEndTime()); err != nil {
		return resp, err
	}
	conn := l.svcCtx.CouponRepository.GetConn()
	tx, err := conn.Begin()
	if err != nil {
		return resp, errors.InternalServerError("tx begin fail")
	}
	err = l.svcCtx.CouponRepository.UseUserCoupon(userCouponId, tx)
	if err != nil {
		tx.Rollback()
		return resp, errors.InternalServerError("update voucher status fail")
	}
	err = l.svcCtx.CouponRepository.InsertUserCouponLog(orderId, userCouponId, tx)
	if err != nil {
		tx.Rollback()
		return resp, errors.InternalServerError("insert log fail")
	}
	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return resp, errors.InternalServerError("tx commit fail")
	}
	return resp, nil
}
