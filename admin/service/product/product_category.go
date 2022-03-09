package product

import (
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
	return model.AddProductCategory(params)
}
func (impl productCategoryServiceImpl) Delete(id uint64) error {
	return model.DeleteProductCategory(id)
}
func (impl productCategoryServiceImpl) Update(params *model.ProductCategory) error {
	return model.UpdateProductCategory(params)
}

func (impl productCategoryServiceImpl) GetById(id uint64) *model.ProductCategory {
	return model.GetProductCategoryById(id)
}
func (impl productCategoryServiceImpl) List(params *model.ProductCategory) []*model.ProductCategory {
	return model.ListProductCategory(params)
}

func (impl productCategoryServiceImpl) Count() int {
	return model.CountProductCategory()
}

func (impl productCategoryServiceImpl) ListPage(pageNum, pageSize int, params *model.ProductCategory) *response.ApiResponse {
	productCategorys, err := model.ListProductCategoryPage(pageNum, pageSize, params)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	total := model.CountProductCategory()
	res := &response.ApiResponse{List: productCategorys, Total: total}
	return res
}
