package logic

import (
	"context"

	. "catering/pkg/log"
	"catering/pkg/utils"
	. "catering/proto/voucher"
	"catering/voucher/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserVoucherLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserVoucherLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserVoucherLogic {
	return &GetUserVoucherLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserVoucherLogic) GetUserVoucher(in *GetUserVoucherReq) (*GetUserVoucherResp, error) {
	resp := &GetUserVoucherResp{}
	var (
		uid   = in.GetUid()
		begin = in.GetBegin()
		count = in.GetCount()
	)
	vouchers := l.svcCtx.VoucherRepository.GetUserVoucher(utils.WithColumn("a.uid", uid))
	Logger.Info(vouchers)
	len := int32(len(vouchers))
	var temp []*UserVoucher
	//如果大于最长
	if begin > len {
		vouchers = []UserVoucher{}
	} else if begin+count > len { //begin小于len但begin+count>len
		vouchers = vouchers[begin:]
	} else { //begin小于len且begin+count<len
		vouchers = vouchers[begin : begin+count]
	}
	Logger.Info(vouchers)
	for index, _ := range vouchers {
		temp = append(temp, &vouchers[index])
	}
	Logger.Info(vouchers)
	resp.Vouchers = temp
	return resp, nil
}
