package logic

import (
	"context"

	"catering/area/pkg/errors"
	"catering/area/rpc/internal/svc"
	pb "catering/proto/common"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetDistinctsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDistinctsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDistinctsLogic {
	return &GetDistinctsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDistinctsLogic) GetDistincts(req *pb.GetDistinctsReq) (*pb.GetDistinctsResp, error) {
	// todo: add your logic here and delete this line
	resp := &pb.GetDistinctsResp{}
	if req.CityId <= 0 {
		return resp, errors.InternalServerError("city_id is not format")
	}
	diss, err := l.svcCtx.DistinctRepository.SelectByCityId(req.CityId)
	if err != nil {
		return resp, err
	}
	resp.Distincts = diss
	return resp, nil
}
