package response

import "catering/model"

type ShopDetail struct {
	model.Shop
	CityName     string `json:"city_name"`
	DistrictName string `json:"district_name"`
	CategoryName string `json:"category_name"`
	ProvinceName string `json:"province_name"`
}
