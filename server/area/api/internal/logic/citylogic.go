package logic

import (
	"context"

	"catering/area/api/internal/svc"
	"catering/area/api/internal/types"
	pb "catering/proto/common"

	"github.com/tal-tech/go-zero/core/logx"
)

type CityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCityLogic(ctx context.Context, svcCtx *svc.ServiceContext) CityLogic {
	return CityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CityLogic) City(req types.GetCitysReq) (*types.GetCitysResp, error) {
	resp := &types.GetCitysResp{}
	res, err := l.svcCtx.CommonRpc.GetCitys(l.ctx, &pb.GetCitysReq{
		ProvinceId: req.ProvinceId,
	})
	if err != nil {
		return resp, err
	}
	resp.Citys = res.GetCitys()
	return resp, nil
}
