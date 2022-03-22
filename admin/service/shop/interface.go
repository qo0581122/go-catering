package shop

import (
	"catering/model"
	"catering/model/common/response"
)

type shopCategoryService interface {
	ListPage(pageNum, pageSize int, params *model.ShopCategory) *response.ApiResponse
	GetOne(params *model.ShopCategory) *model.ShopCategory
	List(params *model.ShopCategory) []*model.ShopCategory
	Count() int
	Add(params *model.ShopCategory) error
	Update(params *model.ShopCategory) error
	Delete(id uint64) error
}

type shopService interface {
	ListPage(pageNum, pageSize int, params *model.Shop) *response.ApiResponse
	GetOne(params *model.Shop) *model.Shop
	List(params *model.Shop) []*model.Shop
	Count() int
	Add(params *model.Shop) error
	Update(params *model.Shop) error
	Delete(id uint64) error
}

type shopProductService interface {
	ListPage(pageNum, pageSize int, params *model.ShopProduct) *response.ApiResponse
	GetOne(params *model.ShopProduct) *model.ShopProduct
	List(params *model.ShopProduct) []*model.ShopProduct
	Count() int
	Add(params *model.ShopProduct) error
	Update(params *model.ShopProduct) error
	Delete(id uint64) error
}
