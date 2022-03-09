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
		err := model.AddProductAttribute(tx, params)
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
			err = model.AddProductAttributeValue(tx, value)
			if err != nil {
				return err
			}
		}
		return nil
	})
}
func (impl productAttributeServiceImpl) Delete(id uint64) error {
	return model.DeleteProductAttribute(id)
}
func (impl productAttributeServiceImpl) Update(params *model.ProductAttribute, values []string) error {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		err := model.UpdateProductAttribute(tx, params)
		if err != nil {
			return err
		}
		id := params.Id
		err = model.DeleteProductAttributeValueByAttributeId(id)
		if err != nil {
			return err
		}
		for _, v := range values {
			value := &model.ProductAttributeValue{
				AttributeId:    id,
				AttributeValue: v,
				Status:         1,
			}
			err = model.AddProductAttributeValue(tx, value)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (impl productAttributeServiceImpl) GetById(id uint64) *model.ProductAttribute {
	return model.GetProductAttributeById(id)
}

type ProductAttributeResponse struct {
	*model.ProductAttribute
	Values []*model.ProductAttributeValue `json:"values"`
}

func (impl productAttributeServiceImpl) List(params *model.ProductAttribute) []*ProductAttributeResponse {
	attributes := model.ListProductAttribute(params)
	attributeValueService := NewProductAttributeValueService()
	var result []*ProductAttributeResponse
	for _, item := range attributes {
		data := attributeValueService.List(&model.ProductAttributeValue{
			AttributeId: item.Id,
		})
		values := make([]*model.ProductAttributeValue, len(data))
		for index, v := range data {
			values[index] = v.ProductAttributeValue
			// values = append(values, v.ProductAttributeValue)
		}
		result = append(result, &ProductAttributeResponse{
			ProductAttribute: item,
			Values:           values,
		})
	}
	return result
}

func (impl productAttributeServiceImpl) Count() int {
	return model.CountProductAttribute()
}

func (impl productAttributeServiceImpl) ListPage(pageNum, pageSize int, params *model.ProductAttribute) *response.ApiResponse {
	productAttributes, err := model.ListProductAttributePage(pageNum, pageSize, params)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	attributeValueService := NewProductAttributeValueService()
	var result []*ProductAttributeResponse
	for _, item := range productAttributes {
		data := attributeValueService.List(&model.ProductAttributeValue{
			AttributeId: item.Id,
		})
		values := make([]*model.ProductAttributeValue, len(data))
		for index, v := range data {
			values[index] = v.ProductAttributeValue
			// values = append(values, v.ProductAttributeValue)
		}
		result = append(result, &ProductAttributeResponse{
			ProductAttribute: item,
			Values:           values,
		})
	}
	total := model.CountProductAttribute()
	res := &response.ApiResponse{List: result, Total: total}
	return res
}
