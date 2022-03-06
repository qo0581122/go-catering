package area

import (
	"catering/model"
	"catering/model/common/response"
	"fmt"
)

var ProvinceService provinceService = NewProvinceService()

func NewProvinceService() provinceService {
	return provinceServiceImpl{}
}

type provinceServiceImpl struct {
}

func (impl provinceServiceImpl) Add(params *model.Province) error {
	return model.AddProvince(params)
}
func (impl provinceServiceImpl) Delete(id uint64) error {
	return model.DeleteProvince(id)
}
func (impl provinceServiceImpl) Update(params *model.Province) error {
	return model.UpdateProvince(params)
}

func (impl provinceServiceImpl) GetById(id uint64) *model.Province {
	return model.GetProvinceById(id)
}
func (impl provinceServiceImpl) List(params *model.Province) []*model.Province {
	return model.ListProvince(params)
}

func (impl provinceServiceImpl) Count() int {
	return model.CountUserAddress()
}

func (impl provinceServiceImpl) ListPage(pageNum, pageSize int, params *model.Province) *response.ApiResponse {
	provinces, err := model.ListProvincePage(pageNum, pageSize, params)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	total := model.CountProvince()
	res := &response.ApiResponse{List: provinces, Total: total}
	return res
}
