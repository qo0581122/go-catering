package shop

import (
	"catering/model"
	"catering/model/common/response"
	"fmt"
)

var ShopService shopService = NewShopService()

type ResponseShop struct {
	*model.Shop
	CityName     string `json:"city_name"`
	DistrictName string `json:"district_name"`
	CategoryName string `json:"category_name"`
	ProvinceName string `json:"province_name"`
}

func NewShopService() shopService {
	return shopServiceImpl{}
}

type shopServiceImpl struct {
}

func (impl shopServiceImpl) Add(params *model.Shop) error {
	province := model.GetProvinceById(params.ProvinceId)
	city := model.GetCityById(params.CityId)
	district := model.GetDistrictById(params.DistrictId)
	params.ShopDetailAddress = province.ProvinceName + city.CityName + district.DistrictName + params.ShopAddress
	return model.AddShop(params)
}
func (impl shopServiceImpl) Delete(id uint64) error {
	return model.DeleteShop(id)
}
func (impl shopServiceImpl) Update(params *model.Shop) error {
	province := model.GetProvinceById(params.ProvinceId)
	city := model.GetCityById(params.CityId)
	district := model.GetDistrictById(params.DistrictId)
	params.ShopDetailAddress = province.ProvinceName + city.CityName + district.DistrictName + params.ShopAddress
	return model.UpdateShop(params)
}

func (impl shopServiceImpl) GetById(id uint64) *model.Shop {
	return model.GetShopById(id)
}
func (impl shopServiceImpl) List(params *model.Shop) []*model.Shop {
	return model.ListShop(params)
}

func (impl shopServiceImpl) Count() int {
	return model.CountUserAddress()
}

func (impl shopServiceImpl) ListPage(pageNum, pageSize int, params *model.Shop) *response.ApiResponse {
	shops, err := model.ListShopPage(pageNum, pageSize, params)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var result []*ResponseShop
	for _, item := range shops {
		province := model.GetProvinceById(item.ProvinceId)
		city := model.GetCityById(item.CityId)
		district := model.GetDistrictById(item.DistrictId)
		category := model.GetShopCategoryById(item.CategoryId)
		result = append(result, &ResponseShop{
			Shop:         item,
			ProvinceName: province.ProvinceName,
			CityName:     city.CityName,
			DistrictName: district.DistrictName,
			CategoryName: category.CategoryName,
		})
	}
	total := model.CountShop()
	res := &response.ApiResponse{List: result, Total: total}
	return res
}
