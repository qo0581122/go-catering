package model

import (
	"catering/global"
	"fmt"

	"gorm.io/gorm"
)

type ProductAttributeValue struct {
	AttributeId    uint64 `json:"attribute_id"`
	AttributeValue string `json:"attribute_value"`
	Status         int    `json:"status"`
	Model
}

func (ProductAttributeValue) TableName() string {
	return "product_attribute_value"
}

func AddProductAttributeValue(tx *gorm.DB, params *ProductAttributeValue) error {
	if tx != nil {
		return tx.Create(&params).Error
	}
	return global.DB.Create(&params).Error
}

func UpdateProductAttributeValue(params *ProductAttributeValue) error {
	return global.DB.Model(&ProductAttributeValue{}).Where("id = ?", params.Id).Updates(&params).Error
}

func DeleteProductAttributeValueById(id uint64) error {
	return global.DB.Delete(&ProductAttributeValue{}, id).Error
}

func DeleteProductAttributeValueByAttributeId(id uint64) error {
	return global.DB.Where("attribute_id", id).Delete(&ProductAttributeValue{}).Error
}

func ListProductAttributeValuePage(pageNum, pageSize int, params *ProductAttributeValue) ([]*ProductAttributeValue, error) {
	var tags []*ProductAttributeValue
	err := global.DB.Where(&params).Scopes(Paginate(pageNum, pageSize)).Find(&tags).Error
	if err != nil {
		return nil, err
	}
	return tags, err
}

func ListProductAttributeValue(params *ProductAttributeValue) []*ProductAttributeValue {
	var tags []*ProductAttributeValue
	err := global.DB.Where(&params).Find(&tags).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return tags
}

func ListProductAttributeValueByAttributeId(id uint64) []*ProductAttributeValue {
	var tags []*ProductAttributeValue
	err := global.DB.Where("attribute_id = ?", id).Find(&tags).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return tags
}

func GetProductAttributeValueById(id uint64) *ProductAttributeValue {
	var res ProductAttributeValue
	err := global.DB.Where("id = ?", id).First(&res).Error
	if err != nil {
		return nil
	}
	return &res
}

func CountProductAttributeValue() int {
	var count int64
	global.DB.Model(&ProductAttributeValue{}).Count(&count)
	return int(count)
}
