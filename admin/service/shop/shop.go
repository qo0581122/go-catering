package shop

import (
	"catering/global"
	"catering/model"
	"catering/model/common/response"
	shopResponse "catering/model/shop/response"
	"errors"
	"fmt"
)

var ShopService shopService = NewShopService()

func NewShopService() shopService {
	return shopServiceImpl{}
}

type shopServiceImpl struct {
}

func (impl shopServiceImpl) Add(params *model.Shop) error {
	province := provinceService.Get(&model.Province{
		Model: model.Model{
			ID: params.ProvinceId,
		},
	})
	city := cityService.Get(&model.City{
		Model: model.Model{
			ID: params.CityId,
		},
	})
	district := districtService.Get(&model.District{
		Model: model.Model{
			ID: params.DistrictId,
		},
	})
	params.ShopDetailAddress = province.ProvinceName + city.CityName + district.DistrictName + params.ShopAddress
	return global.DB.Create(&params).Error
}
func (impl shopServiceImpl) Delete(id uint64) error {
	query := model.ShopProduct{
		ShopId: id,
	}
	var shopProduct model.ShopProduct
	if err := global.DB.Where(&query).Find(&shopProduct).Error; err != nil {
		return err
	}
	if shopProduct.ID > 0 {
		return errors.New("存在关联，无法删除")
	}
	return global.DB.Delete(&model.Shop{}, id).Error
}
func (impl shopServiceImpl) Update(params *model.Shop) error {
	province := provinceService.Get(&model.Province{
		Model: model.Model{
			ID: params.ProvinceId,
		},
	})
	city := cityService.Get(&model.City{
		Model: model.Model{
			ID: params.CityId,
		},
	})
	district := districtService.Get(&model.District{
		Model: model.Model{
			ID: params.DistrictId,
		},
	})
	params.ShopDetailAddress = province.ProvinceName + city.CityName + district.DistrictName + params.ShopAddress
	return global.DB.Model(&model.Shop{}).Where("id = ?", params.ID).Updates(&params).Error
}

func (impl shopServiceImpl) GetOne(params *model.Shop) *model.Shop {
	var res *model.Shop
	err := global.DB.Where(&params).First(&res).Error
	if err != nil {
		return nil
	}
	return res
}

func (impl shopServiceImpl) List(params *model.Shop) []*model.Shop {
	var shops []*model.Shop
	err := global.DB.Where(&params).Find(&shops).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return shops
}

func (impl shopServiceImpl) Count() int {
	var count int64
	global.DB.Model(&model.Shop{}).Count(&count)
	return int(count)
}

func (impl shopServiceImpl) ListPage(pageNum, pageSize int, params *model.Shop) *response.ApiResponse {
	var shops []*model.Shop
	err := global.DB.Preload("Category").Preload("Province").Preload("City").Preload("District").Where(&params).Scopes(model.Paginate(pageNum, pageSize)).Find(&shops).Error
	if err != nil {
		return nil
	}
	var result []*shopResponse.ShopDetail
	for _, item := range shops {
		result = append(result, &shopResponse.ShopDetail{
			Shop:         *item,
			ProvinceName: item.Province.ProvinceName,
			CityName:     item.City.CityName,
			DistrictName: item.District.DistrictName,
			CategoryName: item.Category.CategoryName,
		})
		item.Province = nil
		item.City = nil
		item.Category = nil
		item.District = nil
	}
	var total int64
	global.DB.Model(&model.Shop{}).Count(&total)
	res := &response.ApiResponse{List: result, Total: int(total)}
	return res
}
