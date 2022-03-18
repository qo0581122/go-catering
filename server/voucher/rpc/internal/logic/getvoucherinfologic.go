package logic

import (
	"catering/area/pkg/errors"
	. "catering/pkg/utils"
	. "catering/proto/voucher"
	"catering/voucher/rpc/internal/svc"
	"context"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetVoucherInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVoucherInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVoucherInfoLogic {
	return &GetVoucherInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetVoucherInfoLogic) GetVoucherInfo(in *GetVoucherInfoReq) (*GetVoucherInfoResp, error) {
	resp := &GetVoucherInfoResp{}
	vid := in.GetVoucherId()
	if vid == 0 {
		return nil, errors.InternalServerError("coupon id not zero")
	}
	voucher := l.svcCtx.VoucherRepository.Select(WithColumn("id", vid))
	if len(voucher) > 0 {
		resp.Voucher = &voucher[0]
	}
	return resp, nil
}
