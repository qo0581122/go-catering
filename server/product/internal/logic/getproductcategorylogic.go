package logic

import (
	"context"

	"catering/product/internal/svc"
	pb "catering/proto/product"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetProductCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductCategoryLogic {
	return &GetProductCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProductCategoryLogic) GetProductCategory(req *pb.GetProductCategoryReq) (*pb.GetProductCategoryResp, error) {
	var resp *pb.GetProductCategoryResp
	if req.ProductCategoryId != 0 {
		categorys := l.svcCtx.ProductCategoryRepository.SelectByProductId(req.ProductCategoryId)
		resp.Categorys = []*pb.ProductCategory{categorys}
		return resp, nil
	}
	categorys := l.svcCtx.ProductCategoryRepository.SelectAll()
	resp.Categorys = categorys
	return resp, nil
}
