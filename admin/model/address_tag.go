package model

import (
	"catering/global"
	"fmt"
)

type AddressTag struct {
	Model
	TagName         string `json:"tag_name"`
	TextColor       string `json:"text_color"`
	BorderColor     string `json:"border_color"`
	BackgroundColor string `json:"background_color"`
	Sort            uint32 `json:"sort"`
	Status          uint32 `json:"status"`
}

func (AddressTag) TableName() string {
	return "address_tag"
}

func AddAddressTag(params *AddressTag) error {
	return global.DB.Create(&params).Error
}

func UpdateAddressTag(params *AddressTag) error {
	return global.DB.Model(&AddressTag{}).Where("id = ?", params.Id).Updates(&params).Error
}

func DeleteAddressTag(id uint64) error {
	return global.DB.Delete(&AddressTag{}, id).Error
}

func ListAddressTagPage(pageNum, pageSize int, params *AddressTag) ([]*AddressTag, error) {
	var tags []*AddressTag
	err := global.DB.Where(&params).Scopes(Paginate(pageNum, pageSize)).Find(&tags).Error
	if err != nil {
		return nil, err
	}
	return tags, err
}

func ListAddressTag(params *AddressTag) []*AddressTag {
	var tags []*AddressTag
	err := global.DB.Where(&params).Find(&tags).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return tags
}

func ListAddressTagByAddressId(addressId int) []*AddressTag {
	var tags []*AddressTag
	SQL := "select a.* FROM address_tag a, address_tag_relation b where a.id = b.tag_id AND b.address_id = ? AND a.`status` = 1 AND deleted_time IS NULL ORDER BY sort desc "
	global.DB.Raw(SQL, addressId).Scan(&tags)
	return tags
}

func CountAddressTag() int {
	var count int64
	global.DB.Model(&AddressTag{}).Count(&count)
	return int(count)
}

func GetAddressTagById(id uint64) *AddressTag {
	var res AddressTag
	err := global.DB.Where("id = ?", id).First(&res).Error
	if err != nil {
		return nil
	}
	return &res
}
