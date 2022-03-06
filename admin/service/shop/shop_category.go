package shop

import (
	"catering/model"
	"catering/model/common/response"
	"fmt"
)

var ShopCategoryService shopCategoryService = NewShopCategoryService()

func NewShopCategoryService() shopCategoryService {
	return shopCategoryServiceImpl{}
}

type shopCategoryServiceImpl struct {
}

func (impl shopCategoryServiceImpl) Add(params *model.ShopCategory) error {
	return model.AddShopCategory(params)
}
func (impl shopCategoryServiceImpl) Delete(id uint64) error {
	return model.DeleteShopCategory(id)
}
func (impl shopCategoryServiceImpl) Update(params *model.ShopCategory) error {
	return model.UpdateShopCategory(params)
}

func (impl shopCategoryServiceImpl) GetById(id uint64) *model.ShopCategory {
	return model.GetShopCategoryById(id)
}
func (impl shopCategoryServiceImpl) List(params *model.ShopCategory) []*model.ShopCategory {
	return model.ListShopCategory(params)
}

func (impl shopCategoryServiceImpl) Count() int {
	return model.CountUserAddress()
}

func (impl shopCategoryServiceImpl) ListPage(pageNum, pageSize int, params *model.ShopCategory) *response.ApiResponse {
	shopCategorys, err := model.ListShopCategoryPage(pageNum, pageSize, params)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	total := model.CountShopCategory()
	res := &response.ApiResponse{List: shopCategorys, Total: total}
	return res
}
