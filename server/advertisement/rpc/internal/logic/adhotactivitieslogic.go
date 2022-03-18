package logic

import (
	"context"

	"catering/advertisement/rpc/internal/svc"
	proto "catering/proto/advertisement"

	"github.com/tal-tech/go-zero/core/logx"
)

type AdHotActivitiesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAdHotActivitiesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdHotActivitiesLogic {
	return &AdHotActivitiesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AdHotActivitiesLogic) AdHotActivities(in *proto.GetAdHotActivitieReq) (*proto.GetAdHotActivitieResp, error) {
	var resp proto.GetAdHotActivitieResp
	activities := l.svcCtx.AdHotActivitieRepository.GetList()
	resp.Activities = activities
	return &resp, nil
}
