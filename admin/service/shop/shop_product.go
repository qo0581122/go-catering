package shop

import (
	"catering/global"
	"catering/model"
	"catering/model/common/response"
)

var ShopProductService shopProductService = NewshopProductService()

func NewshopProductService() shopProductService {
	return shopProductServiceImpl{}
}

type shopProductServiceImpl struct {
}

func (impl shopProductServiceImpl) Add(params *model.ShopProduct) error {
	return global.DB.Create(&params).Error
}
func (impl shopProductServiceImpl) Delete(id uint64) error {
	return global.DB.Delete(&model.ShopProduct{}, id).Error
}
func (impl shopProductServiceImpl) Update(params *model.ShopProduct) error {
	return global.DB.Model(&model.ShopProduct{}).Where("id = ?", params.ID).Updates(&params).Error
}

func (impl shopProductServiceImpl) GetOne(params *model.ShopProduct) *model.ShopProduct {
	var res *model.ShopProduct
	err := global.DB.Where(&params).First(&res).Error
	if err != nil {
		return nil
	}
	return res
}

func (impl shopProductServiceImpl) List(params *model.ShopProduct) []*model.ShopProduct {
	var shops []*model.ShopProduct
	err := global.DB.Where(&params).Find(&shops).Error
	if err != nil {
		return nil
	}
	return shops
}

func (impl shopProductServiceImpl) Count() int {
	var count int64
	global.DB.Model(&model.ShopProduct{}).Count(&count)
	return int(count)
}

func (impl shopProductServiceImpl) ListPage(pageNum, pageSize int, params *model.ShopProduct) *response.ApiResponse {
	var shops []*model.ShopProduct
	err := global.DB.Preload("Product").Where(&params).Scopes(model.Paginate(pageNum, pageSize)).Find(&shops).Error
	if err != nil {
		return nil
	}
	var total int64
	global.DB.Model(&model.ShopProduct{}).Count(&total)
	res := &response.ApiResponse{List: shops, Total: int(total)}
	return res
}
