package product

import (
	"catering/model"
	"catering/model/common/response"
	"fmt"
)

var ProductBatchService productBatchService = NewProductBatchService()

func NewProductBatchService() productBatchService {
	return productBatchServiceImpl{}
}

type productBatchServiceImpl struct {
}

func (impl productBatchServiceImpl) Add(params *model.ProductBatch) error {
	return model.AddProductBatch(params)
}
func (impl productBatchServiceImpl) Delete(id uint64) error {
	return model.DeleteProductBatch(id)
}
func (impl productBatchServiceImpl) Update(params *model.ProductBatch) error {
	return model.UpdateProductBatch(params)
}

func (impl productBatchServiceImpl) GetById(id uint64) *model.ProductBatch {
	return model.GetProductBatchById(id)
}
func (impl productBatchServiceImpl) List(params *model.ProductBatch) []*model.ProductBatch {
	return model.ListProductBatch(params)
}

func (impl productBatchServiceImpl) Count() int {
	return model.CountProductBatch()
}

func (impl productBatchServiceImpl) ListPage(pageNum, pageSize int, params *model.ProductBatch) *response.ApiResponse {
	productBatchs, err := model.ListProductBatchPage(pageNum, pageSize, params)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	total := model.CountProductBatch()
	res := &response.ApiResponse{List: productBatchs, Total: total}
	return res
}
