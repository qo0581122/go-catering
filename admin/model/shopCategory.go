package model

import (
	"catering/global"
	"fmt"
)

type ShopCategory struct {
	CategoryName string `json:"category_name"`
	Status       int    `json:"status"`
	Model
}

func (ShopCategory) TableName() string {
	return "shop_category"
}

func AddShopCategory(params *ShopCategory) error {
	return global.DB.Create(&params).Error
}

func UpdateShopCategory(params *ShopCategory) error {
	return global.DB.Model(&ShopCategory{}).Where("id = ?", params.Id).Updates(&params).Error
}

func DeleteShopCategory(id uint64) error {
	return global.DB.Delete(&ShopCategory{}, id).Error
}

func ListShopCategoryPage(pageNum, pageSize int, params *ShopCategory) ([]*ShopCategory, error) {
	var tags []*ShopCategory
	err := global.DB.Where(&params).Scopes(Paginate(pageNum, pageSize)).Find(&tags).Error
	if err != nil {
		return nil, err
	}
	return tags, err
}

func ListShopCategory(params *ShopCategory) []*ShopCategory {
	var tags []*ShopCategory
	err := global.DB.Where(&params).Find(&tags).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return tags
}

func GetShopCategoryById(id uint64) *ShopCategory {
	var res ShopCategory
	err := global.DB.Where("id = ?", id).First(&res).Error
	if err != nil {
		return nil
	}
	return &res
}

func CountShopCategory() int {
	var count int64
	global.DB.Model(&ShopCategory{}).Count(&count)
	return int(count)
}
