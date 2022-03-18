package logic

import (
	"context"

	"catering/area/api/internal/svc"
	"catering/area/api/internal/types"
	pb "catering/proto/common"

	"github.com/tal-tech/go-zero/core/logx"
)

type DistinctLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDistinctLogic(ctx context.Context, svcCtx *svc.ServiceContext) DistinctLogic {
	return DistinctLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DistinctLogic) Distinct(req types.GetDistinctsReq) (*types.GetDistinctsResp, error) {
	resp := &types.GetDistinctsResp{}
	res, err := l.svcCtx.CommonRpc.GetDistincts(l.ctx, &pb.GetDistinctsReq{
		CityId: req.CityId,
	})
	if err != nil {
		return resp, err
	}
	resp.Distincts = res.GetDistincts()
	return resp, nil
}
