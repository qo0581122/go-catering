package product

import (
	"catering/global"
	"catering/model"
	"catering/model/common/response"
	"errors"
	"fmt"
)

var ProductBatchService productBatchService = NewProductBatchService()

func NewProductBatchService() productBatchService {
	return productBatchServiceImpl{}
}

type productBatchServiceImpl struct {
}

func (impl productBatchServiceImpl) Add(params *model.ProductBatch) error {
	return global.DB.Create(&params).Error
}
func (impl productBatchServiceImpl) Delete(id uint64) error {
	query := model.ProductBatchRelation{
		BatchId: id,
	}
	var relation model.ProductBatchRelation
	if err := global.DB.Where(&query).Find(&relation).Error; err != nil {
		return err
	}
	if relation.Id > 0 {
		return errors.New("存在关联，无法删除")
	}
	return global.DB.Delete(&model.ProductBatch{}, id).Error
}
func (impl productBatchServiceImpl) Update(params *model.ProductBatch) error {
	return global.DB.Model(&model.ProductBatch{}).Where("id = ?", params.ID).Updates(&params).Error
}

func (impl productBatchServiceImpl) GetOne(params *model.ProductBatch) *model.ProductBatch {
	var res model.ProductBatch
	err := global.DB.Where(&params).First(&res).Error
	if err != nil {
		return nil
	}
	return &res
}
func (impl productBatchServiceImpl) List(params *model.ProductBatch) []*model.ProductBatch {
	var tags []*model.ProductBatch
	err := global.DB.Where(&params).Find(&tags).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return tags
}

func (impl productBatchServiceImpl) Count() int {
	var count int64
	global.DB.Model(&model.ProductBatch{}).Count(&count)
	return int(count)
}

func (impl productBatchServiceImpl) ListPage(pageNum, pageSize int, params *model.ProductBatch) *response.ApiResponse {
	var productBatchs []*model.ProductBatch
	err := global.DB.Where(&params).Scopes(model.Paginate(pageNum, pageSize)).Find(&productBatchs).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var total int64
	global.DB.Model(&model.ProductBatch{}).Count(&total)
	res := &response.ApiResponse{List: productBatchs, Total: int(total)}
	return res
}
