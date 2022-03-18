package logic

import (
	"context"

	"catering/integration/api/internal/svc"
	. "catering/proto/integration"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserVipUpLogsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserVipUpLogsLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetUserVipUpLogsLogic {
	return GetUserVipUpLogsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserVipUpLogsLogic) GetUserVipUpLogs(req GetUserVipLevelUpLogsReq) (*GetUserVipLevelUpLogsResp, error) {
	return l.svcCtx.IntegrationRpc.GetUserVipLevelUpLogs(l.ctx, &req)
}
