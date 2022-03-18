package logic

import (
	"context"
	"time"

	"catering/area/pkg/errors"
	. "catering/pkg/log"
	. "catering/pkg/utils"
	. "catering/proto/voucher"
	"catering/voucher/rpc/internal/svc"
	. "catering/voucher/rpc/repository"

	"github.com/tal-tech/go-zero/core/logx"
)

type DrawVoucherLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDrawVoucherLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DrawVoucherLogic {
	return &DrawVoucherLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DrawVoucherLogic) DrawVoucher(in *DrawVoucherReq) (*DrawVoucherResp, error) {
	resp := &DrawVoucherResp{}
	var (
		vid   = in.GetVoucherId()
		uid   = in.GetUid()
		count = in.GetCount()
	)
	if vid == 0 || uid == 0 || count == 0 {
		return resp, errors.InternalServerError("vid or uid or count not zero")
	}
	vouchers := l.svcCtx.VoucherRepository.Select(WithColumn("id", vid))
	if len(vouchers) <= 0 {
		return resp, errors.InternalServerError("voucher info is null")
	}
	voucher := vouchers[0]
	Logger.Info(voucher)
	userVouchers := l.svcCtx.VoucherRepository.GetUserVoucher(WithColumn("a.uid", uid), WithColumn("a.voucher_id", vid))
	if len(userVouchers) >= int(voucher.GetCount) {
		return resp, errors.InternalServerError("have been got")
	}
	canGetCount := int(count) - len(userVouchers)
	if err := CheckGetCount(&voucher, canGetCount); err != nil {
		return resp, err
	}
	if err := CheckGetTime(&voucher); err != nil {
		return resp, err
	}
	useBeginTime := voucher.GetUseBeginTime()
	useEndTime := voucher.GetUseEndTime()
	if voucher.GetValidTimeType() == 2 {
		day := int64(voucher.GetValidDay())
		useEndTime = time.Now().Add(time.Duration(day * int64(time.Hour) * 24)).Format(string(DATETIME))
		useBeginTime = GetCurrentTime()
	}
	for i := 0; i < canGetCount; i++ {
		tx, err := l.svcCtx.VoucherRepository.GetConn().Begin()
		if err != nil {
			return resp, err
		}
		err = l.svcCtx.VoucherRepository.InsertUserVoucher(uid, vid, useBeginTime, useEndTime, tx)
		if err != nil {
			tx.Rollback()
			return resp, err
		}
		err = l.svcCtx.VoucherRepository.DecreseVoucherCount(vid, 1, tx)
		if err != nil {
			tx.Rollback()
			return resp, err
		}
		if err = tx.Commit(); err != nil {
			tx.Rollback()
			return resp, err
		}
	}
	return resp, nil
}
