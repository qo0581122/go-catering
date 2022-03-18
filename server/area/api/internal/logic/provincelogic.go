package logic

import (
	"context"

	"catering/area/api/internal/svc"
	"catering/area/api/internal/types"
	pb "catering/proto/common"

	"github.com/tal-tech/go-zero/core/logx"
)

type ProvinceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProvinceLogic(ctx context.Context, svcCtx *svc.ServiceContext) ProvinceLogic {
	return ProvinceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProvinceLogic) Province(req types.GetProvincesReq) (*types.GetProvincesResp, error) {
	resp := &types.GetProvincesResp{}
	res, err := l.svcCtx.CommonRpc.GetProvinces(l.ctx, &pb.GetProvincesReq{})
	if err != nil {
		return resp, err
	}
	resp.Provinces = res.GetProvinces()
	return resp, nil
}
