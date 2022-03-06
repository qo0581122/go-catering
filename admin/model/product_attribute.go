package model

import (
	"catering/global"
	"fmt"

	"gorm.io/gorm"
)

type ProductAttribute struct {
	AttributeName string `json:"attribute_name"`
	Status        int    `json:"status"`
	Model
}

func (ProductAttribute) TableName() string {
	return "product_attribute"
}

func AddProductAttribute(tx *gorm.DB, params *ProductAttribute) error {
	if tx != nil {
		return tx.Create(&params).Error
	}
	return global.DB.Create(&params).Error
}

func UpdateProductAttribute(tx *gorm.DB, params *ProductAttribute) error {
	if tx != nil {
		return tx.Model(&ProductAttribute{}).Where("id = ?", params.Id).Updates(&params).Error
	}
	return global.DB.Model(&ProductAttribute{}).Where("id = ?", params.Id).Updates(&params).Error
}

func DeleteProductAttribute(id uint64) error {
	return global.DB.Delete(&ProductAttribute{}, id).Error
}

func ListProductAttributePage(pageNum, pageSize int, params *ProductAttribute) ([]*ProductAttribute, error) {
	var tags []*ProductAttribute
	err := global.DB.Where(&params).Scopes(Paginate(pageNum, pageSize)).Find(&tags).Error
	if err != nil {
		return nil, err
	}
	return tags, err
}

func ListProductAttribute(params *ProductAttribute) []*ProductAttribute {
	var tags []*ProductAttribute
	err := global.DB.Where(&params).Find(&tags).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return tags
}

func GetProductAttributeById(id uint64) *ProductAttribute {
	var res ProductAttribute
	err := global.DB.Where("id = ?", id).First(&res).Error
	if err != nil {
		return nil
	}
	return &res
}

func CountProductAttribute() int {
	var count int64
	global.DB.Model(&ProductAttribute{}).Count(&count)
	return int(count)
}

func ListProductAttributeByProductId(pid uint64) []*ProductAttribute {
	var res []*ProductAttribute
	SQL := "SELECT a.* FROM product_attribute a, product_attribute_relation b  WHERE a.id = b.attribute_id AND b.product_id = ?"
	global.DB.Raw(SQL, pid).Scan(&res)
	return res
}
