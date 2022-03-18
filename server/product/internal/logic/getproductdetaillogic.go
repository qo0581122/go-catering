package logic

import (
	"context"
	"time"

	"catering/area/pkg/errors"
	"catering/product/internal/svc"
	pb "catering/proto/product"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetProductDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductDetailLogic {
	return &GetProductDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProductDetailLogic) GetProductDetail(req *pb.GetProductDetailReq) (*pb.GetProductDetailResp, error) {
	// todo: add your logic here and delete this line
	var resp *pb.GetProductDetailResp
	if req.ProductId == 0 {
		return resp, errors.BadRequest("1", "product_id is not format")
	}
	productId := req.ProductId
	var productDetail *pb.ProductDetail
	var product *pb.ProductInfo
	i := 4
	baseChan := make(chan *pb.ProductBase)
	attributeChan := make(chan []*pb.ProductAttribute)
	categoryChan := make(chan *pb.ProductCategory)
	batchChan := make(chan []*pb.ProductBatch)
	t, _ := time.ParseDuration("3s")
	timeOut := time.Now().Add(t)
ForEnd:
	for {
		select {
		case baseChan <- l.svcCtx.ProductBaseRepository.SelectByProductIds(productId):
			i -= 1
		case attributeChan <- l.svcCtx.ProductAttributeRepository.SelectByProductId(productId):
			i -= 1
		case categoryChan <- l.svcCtx.ProductCategoryRepository.SelectByProductId(productId):
			i -= 1
		case batchChan <- l.svcCtx.ProductBatchRepository.SelectByProductId(productId):
			i -= 1
		default:
			if i <= 0 || time.Now().After(timeOut) {
				break ForEnd
			}
		}
	}
	productBase := <-baseChan
	productAttribute := <-attributeChan
	productBatch := <-batchChan
	productCategory := <-categoryChan
	product.Base = productBase
	product.Attributes = productAttribute
	product.Batchs = productBatch
	product.Category = productCategory
	productDetail.Product = product
	if productBase.Specis == 2 {
		childProductIds := l.svcCtx.ProductRelationRepository.SelectIdByParentId(productId)
		childProducts := make([]*pb.ProductInfo, len(childProductIds))
		for index, id := range childProductIds {
			go func(index int, id int64, childProducts []*pb.ProductInfo) {
				var childProduct *pb.ProductInfo
				childProductBase := l.svcCtx.ProductBaseRepository.SelectByProductIds(id)
				childProductAttributes := l.svcCtx.ProductAttributeRepository.SelectByProductId(id)
				childProductBatch := l.svcCtx.ProductBatchRepository.SelectByProductId(id)
				childProductCategory := l.svcCtx.ProductCategoryRepository.SelectByProductId(id)
				childProduct.Attributes = childProductAttributes
				childProduct.Base = childProductBase
				childProduct.Batchs = childProductBatch
				childProduct.Category = childProductCategory
				childProducts[index] = childProduct
			}(index, id, childProducts)
		}
		productDetail.ChildProduct = childProducts
	}
	resp.Details = []*pb.ProductDetail{productDetail}
	return resp, nil
}
