package logic

import (
	"context"

	"catering/area/pkg/errors"
	"catering/pkg/utils"
	. "catering/proto/voucher"
	"catering/voucher/rpc/internal/svc"
	. "catering/voucher/rpc/repository"

	"github.com/tal-tech/go-zero/core/logx"
)

type UseVoucherLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUseVoucherLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UseVoucherLogic {
	return &UseVoucherLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UseVoucherLogic) UseVoucher(in *UseVoucherReq) (*UseVoucherResp, error) {
	resp := &UseVoucherResp{}
	var (
		userVoucherId = in.GetUserVoucherId()
		uid           = in.GetUid()
		orderId       = in.GetOrderId()
	)
	if userVoucherId == 0 || uid == 0 || orderId == 0 {
		return resp, errors.InternalServerError("id no zero")
	}
	userVouchers := l.svcCtx.VoucherRepository.GetUserVoucher(utils.WithColumn("a.id", userVoucherId))
	userVoucher := userVouchers[0]
	if err := CheckUseStatus(userVoucher.UseStatus); err != nil {
		return resp, err
	}
	if err := CheckUseTime(userVoucher.GetUseBeginTime(), userVoucher.GetUseEndTime()); err != nil {
		return resp, err
	}
	conn := l.svcCtx.VoucherRepository.GetConn()
	tx, err := conn.Begin()
	if err != nil {
		return resp, errors.InternalServerError("tx begin fail")
	}
	err = l.svcCtx.VoucherRepository.UseUserVoucher(userVoucherId, tx)
	if err != nil {
		tx.Rollback()
		return resp, errors.InternalServerError("update voucher status fail")
	}
	err = l.svcCtx.VoucherRepository.InsertUserVoucherLog(orderId, userVoucherId, tx)
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
