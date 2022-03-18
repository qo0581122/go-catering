package logic

import (
	"context"

	"catering/integration/api/internal/svc"
	. "catering/proto/integration"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserIntegrationLogsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserIntegrationLogsLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetUserIntegrationLogsLogic {
	return GetUserIntegrationLogsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserIntegrationLogsLogic) GetUserIntegrationLogs(req GetUserIntegrationChangeLogsReq) (*GetUserIntegrationChangeLogsResp, error) {
	return l.svcCtx.IntegrationRpc.GetUserIntegrationChangeLogs(l.ctx, &req)
}
