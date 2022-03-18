package logic

import (
	"context"

	. "catering/proto/voucher"
	"catering/voucher/api/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetVoucherInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetVoucherInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetVoucherInfoLogic {
	return GetVoucherInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetVoucherInfoLogic) GetVoucherInfo(req GetVoucherInfoReq) (*GetVoucherInfoResp, error) {
	return l.svcCtx.VoucherRpc.GetVoucherInfo(l.ctx, &req)
}
