package logic

import (
	"context"

	"catering/advertisement/rpc/internal/svc"
	proto "catering/proto/advertisement"
	productProto "catering/proto/product"

	"github.com/tal-tech/go-zero/core/logx"
)

type AdRecommendDailyProductsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAdRecommendDailyProductsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdRecommendDailyProductsLogic {
	return &AdRecommendDailyProductsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AdRecommendDailyProductsLogic) AdRecommendDailyProducts(in *proto.GetAdRecommendDailyProductReq) (*proto.GetAdRecommendDailyProductResp, error) {
	var resp proto.GetAdRecommendDailyProductResp
	var dailys []*proto.AdRecommendDailyProduct
	recommends := l.svcCtx.AdRecommendDailyRepository.GetList()
	for _, recommend := range recommends {
		productId := recommend.GetProductId()
		productResp, _ := l.svcCtx.ProductService.GetProduct(l.ctx, &productProto.GetProductReq{
			ProductId: productId,
		})
		product := productResp.GetBases()[0]
		dailys = append(dailys, &proto.AdRecommendDailyProduct{
			Product: product,
			Ad:      recommend,
		})
	}
	resp.Dailys = dailys
	return &resp, nil
}
