package logic

import (
	"context"

	"catering/area/rpc/internal/svc"
	pb "catering/proto/common"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetProvincesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProvincesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProvincesLogic {
	return &GetProvincesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProvincesLogic) GetProvinces(in *pb.GetProvincesReq) (*pb.GetProvincesResp, error) {
	// todo: add your logic here and delete this line
	resp := &pb.GetProvincesResp{}
	provinces, err := l.svcCtx.ProvinceRepository.SelectAll()
	if err != nil {
		return resp, err
	}
	resp.Provinces = provinces
	return resp, nil
}
