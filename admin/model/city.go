package model

import (
	"catering/global"
	"fmt"
)

type City struct {
	ProvinceId uint64 `json:"province_id"`
	CityName   string `json:"city_name"`
	Status     int    `json:"status"`
	Model
}

func (City) TableName() string {
	return "city"
}

// func (c *City) BeforeCreate(tx *gorm.DB) error {
// 	c.CreatedTime = Time(time.Now())
// 	c.UpdatedTime = Time(time.Now())
// 	return nil
// }

// func (c *City) BeforeUpdate(tx *gorm.DB) error {
// 	c.CreatedTime = Time(time.Now())
// 	return nil
// }

func AddCity(params *City) error {
	return global.DB.Create(&params).Error
}

func UpdateCity(params *City) error {
	return global.DB.Model(&City{}).Where("id = ?", params.Id).Updates(&params).Error
}

func DeleteCity(id uint64) error {
	return global.DB.Delete(&City{}, id).Error
}

func ListCityPage(pageNum, pageSize int, params *City) ([]*City, error) {
	var cities []*City
	if params.CityName != "" {
		global.DB = global.DB.Where("city_name LIKE ?", "%"+params.CityName+"%")
		params.CityName = ""
	}
	err := global.DB.Where(&params).Scopes(Paginate(pageNum, pageSize)).Find(&cities).Error
	if err != nil {
		return nil, err
	}
	return cities, err
}
func ListCity(params *City) []*City {
	var tags []*City
	err := global.DB.Where(&params).Find(&tags).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return tags
}

func ListCityByProvinceId(pid uint64) []*City {
	var cities []*City
	fmt.Println("province_id", pid)
	err := global.DB.Where("province_id = ?", pid).Find(&cities).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return cities
}

func GetCityById(id uint64) *City {
	var city City
	err := global.DB.Where("id = ?", id).First(&city).Error
	if err != nil {
		return nil
	}
	return &city
}

func CountCity() int {
	var count int64
	global.DB.Model(&City{}).Count(&count)
	return int(count)
}
