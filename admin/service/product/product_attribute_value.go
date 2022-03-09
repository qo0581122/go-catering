package product

import (
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
	return model.AddProductAttributeValue(nil, params)
}
func (impl productAttributeValueServiceImpl) Delete(id uint64) error {
	return model.DeleteProductAttributeValueById(id)
}
func (impl productAttributeValueServiceImpl) Update(params *model.ProductAttributeValue) error {
	return model.UpdateProductAttributeValue(params)
}

func (impl productAttributeValueServiceImpl) GetById(id uint64) *model.ProductAttributeValue {
	return model.GetProductAttributeValueById(id)
}
func (impl productAttributeValueServiceImpl) List(params *model.ProductAttributeValue) []*AttributeValueResponseData {
	productAttributeValues := model.ListProductAttributeValue(params)
	var result []*AttributeValueResponseData
	for _, item := range productAttributeValues {
		at := model.GetProductAttributeById(item.AttributeId)
		result = append(result, &AttributeValueResponseData{
			ProductAttributeValue: item,
			AttributeName:         at.AttributeName,
		})
	}
	return result
}

func (impl productAttributeValueServiceImpl) Count() int {
	return model.CountProductAttributeValue()
}

func (impl productAttributeValueServiceImpl) ListPage(pageNum, pageSize int, params *model.ProductAttributeValue) *response.ApiResponse {
	productAttributeValues, err := model.ListProductAttributeValuePage(pageNum, pageSize, params)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var result []*AttributeValueResponseData
	for _, item := range productAttributeValues {
		at := model.GetProductAttributeById(item.AttributeId)
		result = append(result, &AttributeValueResponseData{
			ProductAttributeValue: item,
			AttributeName:         at.AttributeName,
		})
	}
	total := model.CountProductAttributeValue()
	res := &response.ApiResponse{List: result, Total: total}
	return res
}
