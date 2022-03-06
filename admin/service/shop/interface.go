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
