package model

import (
	"catering/global"
	"fmt"
)

type Product struct {
	ProductName   string          `json:"product_name"`
	Description   string          `json:"description"`
	CategoryId    uint64          `json:"category_id"`
	Specis        int             `json:"specis"`
	Url           string          `json:"url"`
	Status        int             `json:"status"`
	OriginalPrice int             `json:"original_price"`
	PayPrice      int             `json:"pay_price"`
	Discount      int             `json:"discount"`
	Category      ProductCategory `gorm:"foreignKey:CategoryId"`
	Model
}

func (Product) TableName() string {
	return "product"
}

func AddProduct(params *Product) error {
	return global.DB.Create(&params).Error
}

func UpdateProduct(params *Product) error {
	return global.DB.Model(&Product{}).Where("id = ?", params.Id).Updates(&params).Error
}

func DeleteProduct(id uint64) error {
	return global.DB.Delete(&Product{}, id).Error
}

func ListProductPage(pageNum, pageSize int, params *Product) ([]*Product, error) {
	var tags []*Product
	err := global.DB.Where(&params).Scopes(Paginate(pageNum, pageSize)).Find(&tags).Error
	if err != nil {
		return nil, err
	}
	return tags, err
}

func ListProduct(params *Product) []*Product {
	var tags []*Product
	err := global.DB.Where(&params).Find(&tags).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return tags
}

func GetProductById(id uint64) *Product {
	var res Product
	err := global.DB.Preload("Category").Where("id = ?", id).First(&res).Error
	if err != nil {
		return nil
	}
	return &res
}

func CountProduct() int {
	var count int64
	global.DB.Model(&Product{}).Count(&count)
	return int(count)
}

func ListChildProductByParentId(id uint64) []*Product {
	var result []*Product
	SQL := "SELECT a.* FROM product a, product_relation b WHERE b.child_product_id = a.id AND b.parent_product_id = ?"
	global.DB.Raw(SQL, id).Scan(&result)
	return result
}
