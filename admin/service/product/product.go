package product

import (
	"catering/global"
	"catering/model"
	"catering/model/common/response"

	"gorm.io/gorm"
)

var ProductService productService = NewProductService()

func NewProductService() productService {
	return productServiceImpl{}
}

type productServiceImpl struct {
}

func (impl productServiceImpl) Add(params *model.Product, attributeIds []uint64, batchIds []uint64, productIds []uint64) error {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&params).Error; err != nil {
			return err
		}
		productId := params.Id
		if params.Specis == 1 {
			for _, id := range attributeIds {
				attributeRelationModel := &model.ProductAttributeRelation{
					ProductId:   productId,
					AttributeId: id,
				}
				if err := tx.Create(&attributeRelationModel).Error; err != nil {
					return err
				}
			}
			for _, item := range batchIds {
				batchRelationModel := &model.ProductBatchRelation{
					ProductId: productId,
					BatchId:   item,
				}
				if err := tx.Create(&batchRelationModel).Error; err != nil {
					return err
				}
			}
		}

		if params.Specis == 2 {
			for _, item := range productIds {
				productRelationModel := &model.ProductRelation{
					ParentProductId: productId,
					ChildProductId:  item,
				}
				if err := tx.Create(&productRelationModel).Error; err != nil {
					return err
				}
			}
		}
		return nil
	})
}
func (impl productServiceImpl) Delete(id uint64) error {
	return model.DeleteProduct(id)
}
func (impl productServiceImpl) Update(params *model.Product, attributeIds []uint64, batchIds []uint64, productIds []uint64) error {
	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Model(&params).Where("id = ?", params.Id).Updates(&params).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Delete("product_id = ?", params.Id).Delete(&model.ProductAttributeRelation{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Delete("product_id = ?", params.Id).Delete(&model.ProductBatchRelation{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Delete("parent_product_id = ?", params.Id).Delete(&model.ProductRelation{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	productId := params.Id
	for _, id := range attributeIds {
		attributeRelationModel := &model.ProductAttributeRelation{
			ProductId:   productId,
			AttributeId: id,
		}
		if err := tx.Create(&attributeRelationModel).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	for _, item := range batchIds {
		batchRelationModel := &model.ProductBatchRelation{
			ProductId: productId,
			BatchId:   item,
		}
		if err := tx.Create(&batchRelationModel).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if params.Specis == 2 {
		for _, item := range productIds {
			productRelationModel := &model.ProductRelation{
				ParentProductId: productId,
				ChildProductId:  item,
			}
			if err := tx.Create(&productRelationModel).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (impl productServiceImpl) GetById(id uint64) *model.Product {
	return model.GetProductById(id)
}
func (impl productServiceImpl) List(params *model.Product) []*model.Product {
	return model.ListProduct(params)
}

func (impl productServiceImpl) Count() int {
	return model.CountProduct()
}

func (impl productServiceImpl) ListPage(pageNum, pageSize int, params *model.Product) *response.ApiResponse {
	var responseDatas []*ResponseData
	products, err := model.ListProductPage(pageNum, pageSize, params)
	if err != nil {
		return nil
	}
	for _, product := range products {
		productId := product.Id

		productAttribute := model.ListProductAttributeByProductId(productId)
		attributeValueService := NewProductAttributeValueService()
		var result []*ProductAttributeResponse
		for _, item := range productAttribute {
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
		category := model.GetProductCategoryById(product.CategoryId)
		productBatch := model.ListProductBatchByProductId(productId)
		var childProducts []*model.Product
		if product.Specis == 2 {
			childProducts = model.ListChildProductByParentId(product.Id)
		}
		responseDatas = append(responseDatas, &ResponseData{
			Product:          product,
			ProductAttibutes: result,
			ProductBatchs:    productBatch,
			ProductCategory:  category,
			ChildProduct:     childProducts,
		})
	}
	total := model.CountProduct()
	res := &response.ApiResponse{List: responseDatas, Total: total}
	return res
}

type ResponseData struct {
	Product          *model.Product              `json:"product"`
	ProductAttibutes []*ProductAttributeResponse `json:"product_attribute"`
	ProductBatchs    []*model.ProductBatch       `json:"product_batch"`
	ProductCategory  *model.ProductCategory      `json:"product_category"`
	ChildProduct     []*model.Product            `json:"child_product"`
}
