// Code generated by goctl. DO NOT EDIT!
// Source: product.proto

package producthandler

import (
	"context"

	pb "catering/proto/product"

	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	GetProductCategoryReq  = pb.GetProductCategoryReq
	GetProductCategoryResp = pb.GetProductCategoryResp
	GetProductDetailReq    = pb.GetProductDetailReq
	GetProductDetailResp   = pb.GetProductDetailResp
	GetProductListReq      = pb.GetProductListReq
	GetProductListResp     = pb.GetProductListResp
	GetProductReq          = pb.GetProductReq
	GetProductResp         = pb.GetProductResp
	ProductAttribute       = pb.ProductAttribute
	ProductAttributeValue  = pb.ProductAttributeValue
	ProductBase            = pb.ProductBase
	ProductBatch           = pb.ProductBatch
	ProductCategory        = pb.ProductCategory
	ProductDetail          = pb.ProductDetail
	ProductInfo            = pb.ProductInfo

	ProductHandler interface {
		GetProduct(ctx context.Context, in *GetProductReq, opts ...grpc.CallOption) (*GetProductResp, error)
		GetProductCategory(ctx context.Context, in *GetProductCategoryReq, opts ...grpc.CallOption) (*GetProductCategoryResp, error)
		GetProductDetail(ctx context.Context, in *GetProductDetailReq, opts ...grpc.CallOption) (*GetProductDetailResp, error)
		GetProductList(ctx context.Context, in *GetProductListReq, opts ...grpc.CallOption) (*GetProductListResp, error)
	}

	defaultProductHandler struct {
		cli zrpc.Client
	}
)

func NewProductHandler(cli zrpc.Client) ProductHandler {
	return &defaultProductHandler{
		cli: cli,
	}
}

func (m *defaultProductHandler) GetProduct(ctx context.Context, in *GetProductReq, opts ...grpc.CallOption) (*GetProductResp, error) {
	client := pb.NewProductHandlerClient(m.cli.Conn())
	return client.GetProduct(ctx, in, opts...)
}

func (m *defaultProductHandler) GetProductCategory(ctx context.Context, in *GetProductCategoryReq, opts ...grpc.CallOption) (*GetProductCategoryResp, error) {
	client := pb.NewProductHandlerClient(m.cli.Conn())
	return client.GetProductCategory(ctx, in, opts...)
}

func (m *defaultProductHandler) GetProductDetail(ctx context.Context, in *GetProductDetailReq, opts ...grpc.CallOption) (*GetProductDetailResp, error) {
	client := pb.NewProductHandlerClient(m.cli.Conn())
	return client.GetProductDetail(ctx, in, opts...)
}

func (m *defaultProductHandler) GetProductList(ctx context.Context, in *GetProductListReq, opts ...grpc.CallOption) (*GetProductListResp, error) {
	client := pb.NewProductHandlerClient(m.cli.Conn())
	return client.GetProductList(ctx, in, opts...)
}
