package logic

import (
	"context"

	. "catering/proto/voucher"
	"catering/voucher/api/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type DrawVoucherLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDrawVoucherLogic(ctx context.Context, svcCtx *svc.ServiceContext) DrawVoucherLogic {
	return DrawVoucherLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DrawVoucherLogic) DrawVoucher(req DrawVoucherReq) (*DrawVoucherResp, error) {
	return l.svcCtx.VoucherRpc.DrawVoucher(l.ctx, &req)
}
