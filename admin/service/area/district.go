package area

import (
	"catering/global"
	"catering/model"
	"catering/model/area/response"
	common "catering/model/common/response"
)

var DistrictService districtService = NewDistrictService()

func NewDistrictService() districtService {
	return districtServiceImpl{}
}

type districtServiceImpl struct {
}

func (impl districtServiceImpl) Add(params *model.District) error {
	return global.DB.Create(&params).Error
}
func (impl districtServiceImpl) Delete(id uint64) error {
	return global.DB.Delete(&model.District{}, id).Error
}
func (impl districtServiceImpl) Update(params *model.District) error {
	return global.DB.Model(&model.District{}).Where("id = ?", params.ID).Updates(&params).Error
}

func (impl districtServiceImpl) Get(params *model.District) *model.District {
	var district model.District
	err := global.DB.Where(&params).First(&district).Error
	if err != nil {
		return nil
	}
	return &district
}

func (impl districtServiceImpl) List(params *model.District) []*model.District {
	var districts []*model.District
	if params.DistrictName != "" {
		global.DB = global.DB.Where("district_name LIKE ?", "%"+params.DistrictName+"%")
		params.DistrictName = ""
	}
	err := global.DB.Where(&params).Order("id desc").Find(&districts).Error
	if err != nil {
		return nil
	}
	return districts
}

func (impl districtServiceImpl) Count() int {
	var total int64
	global.DB.Model(&model.Province{}).Count(&total)
	return int(total)
}

func (impl districtServiceImpl) ListPage(pageNum, pageSize int, params *model.District) *common.ApiResponse {
	res := &common.ApiResponse{}
	var districts []*response.DistrictDetail
	var datas []*model.District
	if params.DistrictName != "" {
		global.DB = global.DB.Where("district_name LIKE ?", "%"+params.DistrictName+"%")
		params.DistrictName = ""
	}
	err := global.DB.Preload("City").Preload("City.Province").Where(&params).Scopes(model.Paginate(pageNum, pageSize)).Find(&datas).Error
	if err != nil {
		return res
	}
	for _, item := range datas {
		cityName := item.City.CityName
		provinceName := item.City.Province.ProvinceName
		provinceId := item.City.ProvinceId
		item.City = nil
		districts = append(districts, &response.DistrictDetail{
			District:     item,
			ProvinceName: provinceName,
			CityName:     cityName,
			ProvinceId:   provinceId,
		})
	}
	var total int64
	global.DB.Model(&model.Province{}).Where(&params).Count(&total)
	res.List = districts
	res.Total = int(total)
	return res
}
