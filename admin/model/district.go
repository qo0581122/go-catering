package model

import (
	"catering/global"
	"fmt"
)

type District struct {
	CityId       uint64 `json:"city_id"`
	DistrictName string `json:"district_name"`
	Status       int    `json:"status"`
	Model
}

func (District) TableName() string {
	return "district"
}

func AddDistrict(params *District) error {
	return global.DB.Create(&params).Error
}

func UpdateDistrict(params *District) error {
	return global.DB.Model(&District{}).Where("id = ?", params.Id).Updates(&params).Error
}

func DeleteDistrict(id uint64) error {
	return global.DB.Delete(&District{}, id).Error
}

func ListDistrictPage(pageNum, pageSize int, params *District) ([]*District, error) {
	var districts []*District
	if params.DistrictName != "" {
		global.DB = global.DB.Where("district_name LIKE ?", "%"+params.DistrictName+"%")
		params.DistrictName = ""
	}
	err := global.DB.Where(&params).Scopes(Paginate(pageNum, pageSize)).Find(&districts).Error
	if err != nil {
		return nil, err
	}
	return districts, err
}

func ListDistrict(params *District) []*District {
	var tags []*District
	err := global.DB.Where(&params).Find(&tags).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return tags
}

func ListByCityId(cid uint64) []*District {
	var districts []*District
	err := global.DB.Where("city_id = ?", cid).Find(&districts).Error
	if err != nil {
		return nil
	}
	return districts
}

func GetDistrictById(id uint64) *District {
	var district District
	err := global.DB.Where("id = ?", id).First(&district).Error
	if err != nil {
		return nil
	}
	return &district
}

func CountDistrict() int {
	var count int64
	global.DB.Model(&District{}).Count(&count)
	return int(count)
}
