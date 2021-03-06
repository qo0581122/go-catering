package area

import (
	"catering/global"
	"catering/model"
	common "catering/model/common/response"
	"errors"
	"fmt"
)

var CityService cityService = NewCityService()

func NewCityService() cityService {
	return cityServiceImpl{}
}

type cityServiceImpl struct {
}

func (impl cityServiceImpl) Add(params *model.City) error {
	return global.DB.Create(&params).Error
}
func (impl cityServiceImpl) Delete(id uint64) error {
	query := model.District{
		CityId: id,
	}
	var district model.District
	if err := global.DB.Where(&query).Find(&district).Error; err != nil {
		return err
	}
	if district.ID > 0 {
		return errors.New("存在关联，无法删除")
	}
	return global.DB.Delete(&model.City{}, id).Error
}
func (impl cityServiceImpl) Update(params *model.City) error {
	return global.DB.Model(&model.City{}).Where("id = ?", params.ID).Updates(&params).Error
}

func (impl cityServiceImpl) Get(params *model.City) *model.City {
	var city model.City
	err := global.DB.Where(&params).First(&city).Error
	if err != nil {
		return nil
	}
	return &city
}

func (impl cityServiceImpl) List(params *model.City) []*model.City {
	var citys []*model.City
	err := global.DB.Where(&params).Find(&citys).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return citys
}

func (impl cityServiceImpl) Count() int {
	var count int64
	global.DB.Model(&model.City{}).Count(&count)
	return int(count)
}

func (impl cityServiceImpl) ListPage(pageNum, pageSize int, params *model.City) *common.ApiResponse {
	res := &common.ApiResponse{}
	var cities []*model.City
	if params.CityName != "" {
		global.DB = global.DB.Where("city_name LIKE ?", "%"+params.CityName+"%")
		params.CityName = ""
	}
	err := global.DB.Preload("Province").Where(&params).Scopes(model.Paginate(pageNum, pageSize)).Order("id desc").Find(&cities).Error
	if err != nil {
		return res
	}
	var count int64
	global.DB.Model(&model.City{}).Where(&params).Count(&count)
	res.List = cities
	res.Total = int(count)
	return res
}
