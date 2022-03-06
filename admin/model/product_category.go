package model

import (
	"catering/global"
	"fmt"
)

type ProductCategory struct {
	CategoryName string `json:"category_name"`
	Status       int    `json:"status"`
	Model
}

func (ProductCategory) TableName() string {
	return "product_category"
}

func AddProductCategory(params *ProductCategory) error {
	return global.DB.Create(&params).Error
}

func UpdateProductCategory(params *ProductCategory) error {
	return global.DB.Model(&ProductCategory{}).Where("id = ?", params.Id).Updates(&params).Error
}

func DeleteProductCategory(id uint64) error {
	return global.DB.Delete(&ProductCategory{}, id).Error
}

func ListProductCategoryPage(pageNum, pageSize int, params *ProductCategory) ([]*ProductCategory, error) {
	var tags []*ProductCategory
	err := global.DB.Where(&params).Scopes(Paginate(pageNum, pageSize)).Find(&tags).Error
	if err != nil {
		return nil, err
	}
	return tags, err
}

func ListProductCategory(params *ProductCategory) []*ProductCategory {
	var tags []*ProductCategory
	err := global.DB.Where(&params).Find(&tags).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return tags
}

func GetProductCategoryById(id uint64) *ProductCategory {
	var res ProductCategory
	err := global.DB.Where("id = ?", id).First(&res).Error
	if err != nil {
		return nil
	}
	return &res
}

func CountProductCategory() int {
	var count int64
	global.DB.Model(&ProductCategory{}).Count(&count)
	return int(count)
}
