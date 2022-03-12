package product

import (
	"catering/global"
	"catering/model"
	"catering/model/common/response"
	"fmt"

	"gorm.io/gorm"
)

var ProductAttributeService productAttributeService = NewProductAttributeService()

func NewProductAttributeService() productAttributeService {
	return productAttributeServiceImpl{}
}

type productAttributeServiceImpl struct {
}

func (impl productAttributeServiceImpl) Add(params *model.ProductAttribute, values []string) error {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&params).Error
		if err != nil {
			return err
		}
		id := params.Id
		for _, v := range values {
			value := &model.ProductAttributeValue{
				AttributeId:    id,
				AttributeValue: v,
				Status:         1,
			}
			err = tx.Create(&value).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
}
func (impl productAttributeServiceImpl) Delete(id uint64) error {
	return global.DB.Delete(&model.ProductAttribute{}, id).Error
}
func (impl productAttributeServiceImpl) Update(params *model.ProductAttribute, values []string) error {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&model.ProductAttribute{}).Where("id = ?", params.Id).Updates(&params).Error
		if err != nil {
			return err
		}
		id := params.Id
		err = global.DB.Where("attribute_id", id).Delete(&model.ProductAttributeValue{}).Error
		if err != nil {
			return err
		}
		for _, v := range values {
			value := &model.ProductAttributeValue{
				AttributeId:    id,
				AttributeValue: v,
				Status:         1,
			}
			err = tx.Create(&value).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (impl productAttributeServiceImpl) GetOne(params *model.ProductAttribute) *model.ProductAttribute {
	var res model.ProductAttribute
	err := global.DB.Where(params).First(&res).Error
	if err != nil {
		return nil
	}
	return &res
}

func (impl productAttributeServiceImpl) List(params *model.ProductAttribute) []*model.ProductAttribute {
	var attributes []*model.ProductAttribute
	err := global.DB.Where(&params).Preload("Values").Find(&attributes).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return attributes
}

func (impl productAttributeServiceImpl) Count() int {
	var count int64
	global.DB.Model(&model.ProductAttribute{}).Count(&count)
	return int(count)
}

func (impl productAttributeServiceImpl) ListPage(pageNum, pageSize int, params *model.ProductAttribute) *response.ApiResponse {
	var attributes []*model.ProductAttribute
	err := global.DB.Where(&params).Scopes(model.Paginate(pageNum, pageSize)).Preload("Values").Find(&attributes).Error
	if err != nil {
		return nil
	}
	var total int64
	global.DB.Model(&model.ProductAttribute{}).Where(params).Count(&total)
	res := &response.ApiResponse{List: attributes, Total: int(total)}
	return res
}
