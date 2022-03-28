package product

import (
	"catering/global"
	"catering/model"
	"catering/model/common/response"
	"fmt"
)

var ProductAttributeValueService productAttributeValueService = NewProductAttributeValueService()

func NewProductAttributeValueService() productAttributeValueService {
	return productAttributeValueServiceImpl{}
}

type AttributeValueResponseData struct {
	*model.ProductAttributeValue `json:",inline"`
	AttributeName                string `json:"attribute_name" form:"attribute_name"`
}

type productAttributeValueServiceImpl struct {
}

func (impl productAttributeValueServiceImpl) Add(params *model.ProductAttributeValue) error {
	return global.DB.Create(&params).Error
}
func (impl productAttributeValueServiceImpl) Delete(id uint64) error {
	return global.DB.Delete(&model.ProductAttributeValue{}, id).Error
}
func (impl productAttributeValueServiceImpl) Update(params *model.ProductAttributeValue) error {
	return global.DB.Model(&model.ProductAttributeValue{}).Where("id = ?", params.ID).Updates(&params).Error
}

func (impl productAttributeValueServiceImpl) GetOne(params *model.ProductAttributeValue) *model.ProductAttributeValue {
	var res model.ProductAttributeValue
	err := global.DB.Where(params).First(&res).Error
	if err != nil {
		return nil
	}
	return &res
}
func (impl productAttributeValueServiceImpl) List(params *model.ProductAttributeValue) []*model.ProductAttributeValue {
	var productAttributeValues []*model.ProductAttributeValue
	err := global.DB.Where(&params).Preload("Attribute").Find(&productAttributeValues).Error
	if err != nil {
		return nil
	}
	return productAttributeValues
}

func (impl productAttributeValueServiceImpl) Count() int {
	var count int64
	global.DB.Model(&model.ProductAttributeValue{}).Count(&count)
	return int(count)
}

func (impl productAttributeValueServiceImpl) ListPage(pageNum, pageSize int, params *model.ProductAttributeValue) *response.ApiResponse {
	var productAttributeValues []*model.ProductAttributeValue
	err := global.DB.Where(&params).Scopes(model.Paginate(pageNum, pageSize)).Preload("Attribute").Find(&productAttributeValues).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var total int64
	global.DB.Model(&model.ProductAttributeValue{}).Where(&params).Count(&total)
	res := &response.ApiResponse{List: productAttributeValues, Total: int(total)}
	return res
}
