package model

import (
	"catering/global"
	"fmt"
)

type ProductBatch struct {
	BatchName string `json:"batch_name"`
	Status    int    `json:"status"`
	Model
}

func (ProductBatch) TableName() string {
	return "product_batch"
}

func AddProductBatch(params *ProductBatch) error {
	return global.DB.Create(&params).Error
}

func UpdateProductBatch(params *ProductBatch) error {
	return global.DB.Model(&ProductBatch{}).Where("id = ?", params.Id).Updates(&params).Error
}

func DeleteProductBatch(id uint64) error {
	return global.DB.Delete(&ProductBatch{}, id).Error
}

func ListProductBatchPage(pageNum, pageSize int, params *ProductBatch) ([]*ProductBatch, error) {
	var tags []*ProductBatch
	err := global.DB.Where(&params).Scopes(Paginate(pageNum, pageSize)).Find(&tags).Error
	if err != nil {
		return nil, err
	}
	return tags, err
}

func ListProductBatch(params *ProductBatch) []*ProductBatch {
	var tags []*ProductBatch
	err := global.DB.Where(&params).Find(&tags).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return tags
}

func GetProductBatchById(id uint64) *ProductBatch {
	var res ProductBatch
	err := global.DB.Where("id = ?", id).First(&res).Error
	if err != nil {
		return nil
	}
	return &res
}

func CountProductBatch() int {
	var count int64
	global.DB.Model(&ProductBatch{}).Count(&count)
	return int(count)
}

func ListProductBatchByProductId(pid uint64) []*ProductBatch {
	var res []*ProductBatch
	SQL := "SELECT a.* FROM product_batch a, product_batch_relation b  WHERE a.id = b.batch_id AND b.product_id = ?"
	global.DB.Raw(SQL, pid).Scan(&res)
	return res
}
