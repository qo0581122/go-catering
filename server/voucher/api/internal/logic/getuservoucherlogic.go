package logic

import (
	"context"

	. "catering/proto/voucher"
	"catering/voucher/api/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserVoucherLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserVoucherLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetUserVoucherLogic {
	return GetUserVoucherLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserVoucherLogic) GetUserVoucher(req GetUserVoucherReq) (*GetUserVoucherResp, error) {
	return l.svcCtx.VoucherRpc.GetUserVoucher(l.ctx, &req)
}
