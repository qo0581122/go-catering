package area

import (
	"catering/model"
	"catering/model/common/response"
	"fmt"
)

type CityDetail struct {
	*model.City
	ProvinceName string `json:"province_name"`
}

var CityService cityService = NewCityService()

func NewCityService() cityService {
	return cityServiceImpl{}
}

type cityServiceImpl struct {
}

func (impl cityServiceImpl) Add(params *model.City) error {
	return model.AddCity(params)
}
func (impl cityServiceImpl) Delete(id uint64) error {
	return model.DeleteCity(id)
}
func (impl cityServiceImpl) Update(params *model.City) error {
	return model.UpdateCity(params)
}

func (impl cityServiceImpl) GetById(id uint64) *model.City {
	return model.GetCityById(id)
}
func (impl cityServiceImpl) List(params *model.City) []*model.City {
	return model.ListCity(params)
}

func (impl cityServiceImpl) Count() int {
	return model.CountUserAddress()
}

func (impl cityServiceImpl) ListPage(pageNum, pageSize int, params *model.City) *response.ApiResponse {
	citys, err := model.ListCityPage(pageNum, pageSize, params)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	total := model.CountCity()
	var detail []*CityDetail
	for _, item := range citys {
		province := model.GetProvinceById(item.ProvinceId)
		detail = append(detail, &CityDetail{
			City:         item,
			ProvinceName: province.ProvinceName,
		})
	}
	res := &response.ApiResponse{List: detail, Total: total}
	return res
}
