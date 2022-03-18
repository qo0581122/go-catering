package logic

import (
	"context"

	"catering/area/pkg/errors"
	"catering/integration/rpc/internal/svc"
	. "catering/proto/integration"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserVipInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserVipInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserVipInfoLogic {
	return &GetUserVipInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserVipInfoLogic) GetUserVipInfo(in *GetUserVipInfoReq) (*GetUserVipInfoResp, error) {
	resp := &GetUserVipInfoResp{}
	uid := in.GetUid()
	if uid == 0 {
		return resp, errors.InternalServerError("uid is null")
	}
	integration := l.svcCtx.IntegrationRepository.GetUserIntegration(uid)
	vip := l.svcCtx.IntegrationRepository.GetUserVipLevel(uid)
	resp.Integration = integration
	resp.Level = vip
	return resp, nil
}
