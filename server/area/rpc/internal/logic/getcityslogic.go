package logic

import (
	"context"

	"catering/area/pkg/errors"
	"catering/area/rpc/internal/svc"
	pb "catering/proto/common"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetCitysLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCitysLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCitysLogic {
	return &GetCitysLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCitysLogic) GetCitys(req *pb.GetCitysReq) (*pb.GetCitysResp, error) {
	// todo: add your logic here and delete this line
	resp := &pb.GetCitysResp{}
	if req.ProvinceId <= 0 {
		return resp, errors.InternalServerError("province_id is not format")
	}
	citys, err := l.svcCtx.CityRepository.SelectByProvinceId(req.ProvinceId)
	if err != nil {
		return resp, err
	}
	resp.Citys = citys
	return resp, nil
}
