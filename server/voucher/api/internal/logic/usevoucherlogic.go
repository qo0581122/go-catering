package logic

import (
	"context"

	. "catering/proto/voucher"
	"catering/voucher/api/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type UseVoucherLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUseVoucherLogic(ctx context.Context, svcCtx *svc.ServiceContext) UseVoucherLogic {
	return UseVoucherLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UseVoucherLogic) UseVoucher(req UseVoucherReq) (*UseVoucherResp, error) {
	return l.svcCtx.VoucherRpc.UseVoucher(l.ctx, &req)
}
