package logic

import (
	"context"

	"catering/area/pkg/errors"
	"catering/coupon/rpc/internal/svc"
	. "catering/coupon/rpc/repository"
	. "catering/proto/coupon"

	"github.com/tal-tech/go-zero/core/logx"
)

type DrawCouponLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDrawCouponLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DrawCouponLogic {
	return &DrawCouponLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DrawCouponLogic) DrawCoupon(in *DrawCouponReq) (*DrawCouponResp, error) {
	resp := &DrawCouponResp{}
	var (
		cid     = in.GetCouponId()
		uid     = in.GetUid()
		count   = in.GetCount()
		getType = in.GetGetType()
	)
	if cid == 0 || uid == 0 || count == 0 {
		return resp, errors.InternalServerError("cid or uid or count not zero")
	}
	counpon := l.svcCtx.CouponRepository.GetCouponInfo(cid)
	if counpon == nil {
		return resp, errors.InternalServerError("counpon info is null")
	}
	userCoupons := l.svcCtx.CouponRepository.GetUserCouponsByCid(uid, cid)
	if len(userCoupons) >= int(counpon.GetCount) {
		return resp, errors.InternalServerError("have been got")
	}
	canGetCount := int(count) - len(userCoupons)
	if err := CheckGetCount(counpon, canGetCount); err != nil {
		return resp, err
	}
	if err := CheckGetTime(counpon); err != nil {
		return resp, err
	}
	useBeginTime, useEndTime := GetUseTime(counpon)
	for i := 0; i < canGetCount; i++ {
		tx, err := l.svcCtx.CouponRepository.GetConn().Begin()
		if err != nil {
			continue
		}
		err = l.svcCtx.CouponRepository.InserUserCoupon(uid, cid, useBeginTime, useEndTime, getType, tx)
		if err != nil {
			tx.Rollback()
			continue
		}
		err = l.svcCtx.CouponRepository.DecreseCouponCount(cid, 1, tx)
		if err != nil {
			tx.Rollback()
			continue
		}
		if err = tx.Commit(); err != nil {
			tx.Rollback()
			continue
		}
	}
	return resp, nil
}
