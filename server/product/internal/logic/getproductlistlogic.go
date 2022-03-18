package logic

import (
	"catering/product/internal/svc"
	pb "catering/proto/product"
	"context"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetProductListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductListLogic {
	return &GetProductListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProductListLogic) GetProductList(in *pb.GetProductListReq) (*pb.GetProductListResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetProductListResp{}, nil
}
