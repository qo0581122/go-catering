package product

import (
	"catering/model"
	"catering/model/common/response"
)

type productService interface {
	ListPage(pageNum, pageSize int, params *model.Product) *response.ApiResponse
	GetOne(params *model.Product) *model.Product
	List(params *model.Product) []*model.Product
	Count() int
	Add(params *model.Product, attributeIds []uint64, batchIds []uint64, productIds []uint64) error
	Update(params *model.Product, attributeIds []uint64, batchIds []uint64, productIds []uint64) error
	Delete(id uint64) error
}

type productCategoryService interface {
	ListPage(pageNum, pageSize int, params *model.ProductCategory) *response.ApiResponse
	GetOne(params *model.ProductCategory) *model.ProductCategory
	List(params *model.ProductCategory) []*model.ProductCategory
	Count() int
	Add(params *model.ProductCategory) error
	Update(params *model.ProductCategory) error
	Delete(id uint64) error
}

type productBatchService interface {
	ListPage(pageNum, pageSize int, params *model.ProductBatch) *response.ApiResponse
	GetOne(params *model.ProductBatch) *model.ProductBatch
	List(params *model.ProductBatch) []*model.ProductBatch
	Count() int
	Add(params *model.ProductBatch) error
	Update(params *model.ProductBatch) error
	Delete(id uint64) error
}

type productAttributeService interface {
	ListPage(pageNum, pageSize int, params *model.ProductAttribute) *response.ApiResponse
	GetOne(params *model.ProductAttribute) *model.ProductAttribute
	List(params *model.ProductAttribute) []*model.ProductAttribute
	Count() int
	Add(params *model.ProductAttribute, values []string) error
	Update(params *model.ProductAttribute, values []string) error
	Delete(id uint64) error
}

type productAttributeValueService interface {
	ListPage(pageNum, pageSize int, params *model.ProductAttributeValue) *response.ApiResponse
	GetOne(params *model.ProductAttributeValue) *model.ProductAttributeValue
	List(params *model.ProductAttributeValue) []*model.ProductAttributeValue
	Count() int
	Add(params *model.ProductAttributeValue) error
	Update(params *model.ProductAttributeValue) error
	Delete(id uint64) error
}
