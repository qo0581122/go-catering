package model

import (
	"catering/global"
	"fmt"
)

type UserAddress struct {
	Model
	Uid           uint64 `json:"uid"`
	ReceiveName   string `json:"receive_name"`
	ReceiveSex    uint32 `json:"receive_sex"`
	ReceivePhone  string `json:"receive_phone"`
	ProvinceId    uint64 `json:"province_id"`
	CityId        uint64 `json:"city_id"`
	DistinctId    uint64 `json:"distinct_id"`
	DetailAddress string `json:"detail_address"`
	IsDefault     uint32 `json:"is_default"`
	Sort          uint32 `json:"sort"`
}

func (UserAddress) TableName() string {
	return "user_address"
}

func AddUserAddress(params *UserAddress) error {
	return global.DB.Create(&params).Error
}

func UpdateUserAddress(params *UserAddress) error {
	return global.DB.Model(&UserAddress{}).Where("id = ?", params.Id).Updates(&params).Error
}

func DeleteUserAddress(id uint64) error {
	return global.DB.Delete(&UserAddress{}, id).Error
}

func ListUserAddressPage(pageNum, pageSize int, params *UserAddress) ([]*UserAddress, error) {
	var tags []*UserAddress
	err := global.DB.Where(&params).Scopes(Paginate(pageNum, pageSize)).Find(&tags).Error
	if err != nil {
		return nil, err
	}
	return tags, err
}

func ListUserAddress(params *UserAddress) []*UserAddress {
	var tags []*UserAddress
	err := global.DB.Where(&params).Find(&tags).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return tags
}

func GetUserAddressById(id uint64) *UserAddress {
	var res UserAddress
	err := global.DB.Where("id = ?", id).First(&res).Error
	if err != nil {
		return nil
	}
	return &res
}

func CountUserAddress() int {
	var count int64
	global.DB.Model(&UserAddress{}).Count(&count)
	return int(count)
}
