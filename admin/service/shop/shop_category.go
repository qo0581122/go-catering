package shop

import (
	"catering/global"
	"catering/model"
	"catering/model/common/response"
	"errors"
	"fmt"
)

var ShopCategoryService shopCategoryService = NewShopCategoryService()

func NewShopCategoryService() shopCategoryService {
	return shopCategoryServiceImpl{}
}

type shopCategoryServiceImpl struct {
}

func (impl shopCategoryServiceImpl) Add(params *model.ShopCategory) error {
	return global.DB.Create(&params).Error
}
func (impl shopCategoryServiceImpl) Delete(id uint64) error {
	query := model.Shop{
		CategoryId: id,
	}
	var shop model.Shop
	if err := global.DB.Where(&query).Find(&shop).Error; err != nil {
		return err
	}
	if shop.ID > 0 {
		return errors.New("存在关联，无法删除")
	}
	return global.DB.Delete(&model.ShopCategory{}, id).Error
}
func (impl shopCategoryServiceImpl) Update(params *model.ShopCategory) error {
	return global.DB.Model(&model.ShopCategory{}).Where("id = ?", params.ID).Updates(&params).Error
}

func (impl shopCategoryServiceImpl) GetOne(params *model.ShopCategory) *model.ShopCategory {
	var res model.ShopCategory
	err := global.DB.Where(&params).First(&res).Error
	if err != nil {
		return nil
	}
	return &res
}
func (impl shopCategoryServiceImpl) List(params *model.ShopCategory) []*model.ShopCategory {
	var res []*model.ShopCategory
	err := global.DB.Where(&params).Find(&res).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return res
}

func (impl shopCategoryServiceImpl) Count() int {
	var total int64
	global.DB.Model(&model.ShopCategory{}).Count(&total)
	return int(total)
}

func (impl shopCategoryServiceImpl) ListPage(pageNum, pageSize int, params *model.ShopCategory) *response.ApiResponse {
	var shopCategorys []*model.ShopCategory
	err := global.DB.Where(&params).Find(&shopCategorys).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var total int64
	global.DB.Model(&model.ShopCategory{}).Where(&params).Count(&total)
	return &response.ApiResponse{List: shopCategorys, Total: int(total)}
}
