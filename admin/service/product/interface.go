package product

import (
	"catering/model"
	"catering/model/common/response"
)

type productService interface {
	ListPage(pageNum, pageSize int, params *model.Product) *response.ApiResponse
	GetById(id uint64) *model.Product
	List(params *model.Product) []*model.Product
	Count() int
	Add(params *model.Product, attributeIds []uint64, batchIds []uint64, productIds []uint64) error
	Update(params *model.Product, attributeIds []uint64, batchIds []uint64, productIds []uint64) error
	Delete(id uint64) error
}

type productCategoryService interface {
	ListPage(pageNum, pageSize int, params *model.ProductCategory) *response.ApiResponse
	GetById(id uint64) *model.ProductCategory
	List(params *model.ProductCategory) []*model.ProductCategory
	Count() int
	Add(params *model.ProductCategory) error
	Update(params *model.ProductCategory) error
	Delete(id uint64) error
}

type productBatchService interface {
	ListPage(pageNum, pageSize int, params *model.ProductBatch) *response.ApiResponse
	GetById(id uint64) *model.ProductBatch
	List(params *model.ProductBatch) []*model.ProductBatch
	Count() int
	Add(params *model.ProductBatch) error
	Update(params *model.ProductBatch) error
	Delete(id uint64) error
}

type productAttributeService interface {
	ListPage(pageNum, pageSize int, params *model.ProductAttribute) *response.ApiResponse
	GetById(id uint64) *model.ProductAttribute
	List(params *model.ProductAttribute) []*ProductAttributeResponse
	Count() int
	Add(params *model.ProductAttribute, values []string) error
	Update(params *model.ProductAttribute, values []string) error
	Delete(id uint64) error
}

type productAttributeValueService interface {
	ListPage(pageNum, pageSize int, params *model.ProductAttributeValue) *response.ApiResponse
	GetById(id uint64) *model.ProductAttributeValue
	List(params *model.ProductAttributeValue) []*AttributeValueResponseData
	Count() int
	Add(params *model.ProductAttributeValue) error
	Update(params *model.ProductAttributeValue) error
	Delete(id uint64) error
}
