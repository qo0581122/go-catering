package model

import (
	"catering/global"
	"fmt"
)

type Province struct {
	ProvinceName string `json:"province_name" xorm:"not null VARCHAR(255)"`
	Status       int    `json:"status" xorm:"not null default 1 INT(1)"`
	Model
}

func (Province) TableName() string {
	return "province"
}

func AddProvince(params *Province) error {
	return global.DB.Create(&params).Error
}

func UpdateProvince(params *Province) error {
	return global.DB.Model(&Province{}).Where("id = ?", params.Id).Updates(&params).Error
}

func DeleteProvince(id uint64) error {
	return global.DB.Delete(&Province{}, id).Error
}

func ListProvincePage(pageNum, pageSize int, params *Province) ([]*Province, error) {
	var tags []*Province
	if params.ProvinceName != "" {
		global.DB = global.DB.Where("district_name LIKE ?", "%"+params.ProvinceName+"%")
		params.ProvinceName = ""
	}
	err := global.DB.Where(&params).Scopes(Paginate(pageNum, pageSize)).Find(&tags).Error
	if err != nil {
		return nil, err
	}
	return tags, err
}

func ListProvince(params *Province) []*Province {
	var tags []*Province
	err := global.DB.Where(&params).Find(&tags).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return tags
}

func GetProvinceById(id uint64) *Province {
	var res Province
	err := global.DB.Where("id = ?", id).First(&res).Error
	if err != nil {
		return nil
	}
	return &res
}

func CountProvince() int {
	var count int64
	global.DB.Model(&Province{}).Count(&count)
	return int(count)
}
