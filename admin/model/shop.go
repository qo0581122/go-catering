package model

import (
	"catering/global"
	"fmt"
)

type Shop struct {
	CategoryId        uint64 `json:"category_id"`
	ShopName          string `json:"shop_name"`
	ShopAddress       string `json:"shop_address"`
	Longitude         string `json:"logintude"`
	Latitude          string `json:"latitude"`
	BusinessFlag      int    `json:"business_flag"`
	ContactNumber     string `json:"contact_number"`
	Status            int    `json:"status"`
	DistrictId        uint64 `json:"district_id"`
	ShopDetailAddress string `json:"shop_detail_address"`
	ProvinceId        uint64 `json:"province_id"`
	CityId            uint64 `json:"city_id"`
	Model
}

func (Shop) TableName() string {
	return "shop"
}

func AddShop(params *Shop) error {
	return global.DB.Create(&params).Error
}

func UpdateShop(params *Shop) error {
	return global.DB.Model(&Shop{}).Where("id = ?", params.Id).Updates(&params).Error
}

func DeleteShop(id uint64) error {
	return global.DB.Delete(&Shop{}, id).Error
}

func ListShopPage(pageNum, pageSize int, params *Shop) ([]*Shop, error) {
	var tags []*Shop
	err := global.DB.Where(&params).Scopes(Paginate(pageNum, pageSize)).Find(&tags).Error
	if err != nil {
		return nil, err
	}
	return tags, err
}

func ListShop(params *Shop) []*Shop {
	var tags []*Shop
	err := global.DB.Where(&params).Find(&tags).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return tags
}
func CountShop() int {
	var count int64
	global.DB.Model(&Shop{}).Count(&count)
	return int(count)
}

func GetShopById(id uint64) *Shop {
	var res Shop
	err := global.DB.Where("id = ?", id).First(&res).Error
	if err != nil {
		return nil
	}
	return &res
}
