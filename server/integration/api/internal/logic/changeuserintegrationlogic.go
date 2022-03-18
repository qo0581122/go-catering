package logic

import (
	"context"

	"catering/integration/api/internal/svc"
	. "catering/proto/integration"

	"github.com/tal-tech/go-zero/core/logx"
)

type ChangeUserIntegrationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangeUserIntegrationLogic(ctx context.Context, svcCtx *svc.ServiceContext) ChangeUserIntegrationLogic {
	return ChangeUserIntegrationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangeUserIntegrationLogic) ChangeUserIntegration(req ChangeUserIntegrationReq) (*ChangeUserIntegrationResp, error) {
	resp, err := l.svcCtx.IntegrationRpc.ChangeUserIntegration(l.ctx, &req)
	return resp, err
}
