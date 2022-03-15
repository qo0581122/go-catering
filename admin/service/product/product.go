package product

import (
	"catering/global"
	"catering/model"
	"catering/model/common/response"
	productResp "catering/model/product/response"
	"fmt"

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
		productId := params.ID
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
	return global.DB.Delete(&model.Product{}, id).Error
}
func (impl productServiceImpl) Update(params *model.Product, attributeIds []uint64, batchIds []uint64, productIds []uint64) error {
	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Model(&params).Where("id = ?", params.ID).Updates(&params).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Delete("product_id = ?", params.ID).Delete(&model.ProductAttributeRelation{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Delete("product_id = ?", params.ID).Delete(&model.ProductBatchRelation{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Delete("parent_product_id = ?", params.ID).Delete(&model.ProductRelation{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	productId := params.ID
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

func (impl productServiceImpl) GetOne(params *model.Product) *model.Product {
	var res model.Product
	err := global.DB.Preload("Category").Where(&params).First(&res).Error
	if err != nil {
		return nil
	}
	return &res
}
func (impl productServiceImpl) List(params *model.Product) []*model.Product {
	var products []*model.Product
	err := global.DB.Where(&params).Find(&products).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return products
}

func (impl productServiceImpl) Count() int {
	var count int64
	global.DB.Model(&model.Product{}).Count(&count)
	return int(count)
}

func (impl productServiceImpl) ListPage(pageNum, pageSize int, params *model.Product) *response.ApiResponse {
	var responseDatas []*productResp.ResponseData
	var products []*model.Product
	err := global.DB.Where(&params).Scopes(model.Paginate(pageNum, pageSize)).Preload("Category").Preload("Attributes").Preload("Batchs").Find(&products).Error
	if err != nil {
		return nil
	}
	for _, product := range products {
		productId := product.ID
		var childProducts []*model.Product
		if product.Specis == 2 {
			var result []*model.Product
			SQL := "SELECT a.* FROM product a, product_relation b WHERE b.child_product_id = a.id AND b.parent_product_id = ?"
			global.DB.Raw(SQL, productId).Scan(&result)
			childProducts = result
		}
		responseDatas = append(responseDatas, &productResp.ResponseData{
			Product:      product,
			ChildProduct: childProducts,
		})
	}
	var total int64
	global.DB.Model(&model.Product{}).Where(&params).Count(&total)
	res := &response.ApiResponse{List: responseDatas, Total: int(total)}
	return res
}
