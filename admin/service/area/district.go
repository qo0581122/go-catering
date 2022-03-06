package area

import (
	"catering/model"
	"catering/model/common/response"
	"fmt"
)

type DistrictDetail struct {
	*model.District
	ProvinceName string `json:"province_name"`
	CityName     string `json:"city_name"`
	ProvinceId   uint64 `json:"province_id"`
}

var DistrictService districtService = NewDistrictService()

func NewDistrictService() districtService {
	return districtServiceImpl{}
}

type districtServiceImpl struct {
}

func (impl districtServiceImpl) Add(params *model.District) error {
	return model.AddDistrict(params)
}
func (impl districtServiceImpl) Delete(id uint64) error {
	return model.DeleteDistrict(id)
}
func (impl districtServiceImpl) Update(params *model.District) error {
	return model.UpdateDistrict(params)
}

func (impl districtServiceImpl) GetById(id uint64) *model.District {
	return model.GetDistrictById(id)
}
func (impl districtServiceImpl) List(params *model.District) []*model.District {
	return model.ListDistrict(params)
}

func (impl districtServiceImpl) Count() int {
	return model.CountUserAddress()
}

func (impl districtServiceImpl) ListPage(pageNum, pageSize int, params *model.District) *response.ApiResponse {
	districts, err := model.ListDistrictPage(pageNum, pageSize, params)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	total := model.CountDistrict()
	var detail []*DistrictDetail
	for _, item := range districts {
		city := model.GetCityById(item.CityId)
		province := model.GetProvinceById(city.ProvinceId)
		detail = append(detail, &DistrictDetail{
			District:     item,
			ProvinceName: province.ProvinceName,
			CityName:     city.CityName,
			ProvinceId:   province.Id,
		})
	}
	res := &response.ApiResponse{List: detail, Total: total}
	return res
}
