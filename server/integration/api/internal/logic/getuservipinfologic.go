package logic

import (
	"context"

	"catering/integration/api/internal/svc"
	. "catering/proto/integration"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserVipInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserVipInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetUserVipInfoLogic {
	return GetUserVipInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserVipInfoLogic) GetUserVipInfo(req GetUserVipInfoReq) (*GetUserVipInfoResp, error) {
	return l.svcCtx.IntegrationRpc.GetUserVipInfo(l.ctx, &req)
}
