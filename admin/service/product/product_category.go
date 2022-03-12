package product

import (
	"catering/global"
	"catering/model"
	"catering/model/common/response"
	"fmt"
)

var ProductCategoryService productCategoryService = NewProductCategoryService()

func NewProductCategoryService() productCategoryService {
	return productCategoryServiceImpl{}
}

type productCategoryServiceImpl struct {
}

func (impl productCategoryServiceImpl) Add(params *model.ProductCategory) error {
	return global.DB.Create(&params).Error
}
func (impl productCategoryServiceImpl) Delete(id uint64) error {
	return global.DB.Delete(&model.ProductCategory{}, id).Error
}
func (impl productCategoryServiceImpl) Update(params *model.ProductCategory) error {
	return global.DB.Model(&model.ProductCategory{}).Where("id = ?", params.Id).Updates(&params).Error
}

func (impl productCategoryServiceImpl) GetOne(params *model.ProductCategory) *model.ProductCategory {
	var res model.ProductCategory
	err := global.DB.Where(&params).First(&res).Error
	if err != nil {
		return nil
	}
	return &res
}
func (impl productCategoryServiceImpl) List(params *model.ProductCategory) []*model.ProductCategory {
	var tags []*model.ProductCategory
	err := global.DB.Where(&params).Find(&tags).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return tags
}

func (impl productCategoryServiceImpl) Count() int {
	var count int64
	global.DB.Model(&model.ProductCategory{}).Count(&count)
	return int(count)
}

func (impl productCategoryServiceImpl) ListPage(pageNum, pageSize int, params *model.ProductCategory) *response.ApiResponse {
	var productCategorys []*model.ProductCategory
	err := global.DB.Where(&params).Scopes(model.Paginate(pageNum, pageSize)).Find(&productCategorys).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var total int64
	global.DB.Model(&model.ProductCategory{}).Count(&total)
	res := &response.ApiResponse{List: productCategorys, Total: int(total)}
	return res
}
