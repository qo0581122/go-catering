package logic

import (
	"context"

	"catering/advertisement/api/internal/svc"
	"catering/advertisement/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetHotActivitieLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetHotActivitieLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetHotActivitieLogic {
	return GetHotActivitieLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetHotActivitieLogic) GetHotActivitie(req types.GetAdHotActivitieReq) (*types.GetAdHotActivitieResp, error) {
	// todo: add your logic here and delete this line

	return &types.GetAdHotActivitieResp{}, nil
}
