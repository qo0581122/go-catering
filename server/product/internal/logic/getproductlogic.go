package logic

import (
	"context"

	"catering/product/internal/svc"
	pb "catering/proto/product"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductLogic {
	return &GetProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProductLogic) GetProduct(req *pb.GetProductReq) (*pb.GetProductResp, error) {
	var resp *pb.GetProductResp
	if req.ProductId != 0 {
		product := l.svcCtx.ProductBaseRepository.SelectByProductIds(req.ProductId)
		resp.Bases = []*pb.ProductBase{product}
		return resp, nil
	}
	products := l.svcCtx.ProductBaseRepository.SelectAll()
	resp.Bases = products
	return resp, nil
}
