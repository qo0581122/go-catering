package logic

import (
	"context"
	"sync"

	"catering/area/rpc/internal/svc"
	pb "catering/proto/common"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetProvincesDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProvincesDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProvincesDetailLogic {
	return &GetProvincesDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProvincesDetailLogic) GetProvincesDetail(in *pb.GetProvincesDetailReq) (*pb.GetProvincesDetailResp, error) {
	resp := &pb.GetProvincesDetailResp{}
	provinces, err := l.svcCtx.ProvinceRepository.SelectAll()
	if err != nil {
		return resp, err
	}
	var wr sync.WaitGroup
	for i, province := range provinces {
		wr.Add(1)
		go func(index int, province *pb.Province) {
			defer wr.Done()
			provinceId := province.GetId()
			citys, _ := l.svcCtx.CityRepository.SelectByProvinceId(provinceId)
			var wr2 sync.WaitGroup
			for j, city := range citys {
				wr2.Add(1)
				go func(index int, city *pb.City) {
					defer wr2.Done()
					cityId := city.GetId()
					distincts, _ := l.svcCtx.DistinctRepository.SelectByCityId(cityId)
					citys[index].Distincts = distincts
				}(j, city)
			}
			wr2.Wait()
			provinces[index].Citys = citys
		}(i, province)
	}
	wr.Wait()
	resp.Provinces = provinces
	return resp, nil
}
