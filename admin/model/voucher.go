package model

import (
	"catering/global"
	"fmt"
)

type Voucher struct {
	Model
	Name          string `gorm:"column:voucher_name" json:"voucher_name"`
	LeastUsePrice uint64 `json:"least_use_price"`
	Price         uint64 `json:"price"`
	PicUrl        string `json:"pic_url"`
	Description   string `json:"description"`
	TotalCount    uint64 `json:"total_count"`
	RemainCount   uint64 `json:"remain_count"`
	GetType       uint32 `json:"get_type"`
	ValidTimeType uint32 `json:"valid_time_type"`
	ValidDay      uint32 `json:"valid_day"`
	UseType       uint32 `json:"use_type"`
	Status        uint32 `json:"status"`
	GetBeginTime  Time   `json:"get_begin_time"`
	GetEndTime    Time   `json:"get_end_time"`
	UseBeginTime  Time   `json:"use_begin_time"`
	UseEndTime    Time   `json:"use_end_time"`
	GetCount      uint32 `json:"get_count"`
}

func (Voucher) TableName() string {
	return "voucher"
}

func AddVoucher(params *Voucher) error {
	return global.DB.Create(&params).Error
}

func UpdateVoucher(params *Voucher) error {
	return global.DB.Model(&Voucher{}).Where("id = ?", params.Id).Updates(&params).Error
}

func DeleteVoucher(id uint64) error {
	return global.DB.Delete(&Voucher{}, id).Error
}

func ListVoucherPage(pageNum, pageSize int, params *Voucher) ([]*Voucher, error) {
	var tags []*Voucher
	err := global.DB.Where(&params).Scopes(Paginate(pageNum, pageSize)).Find(&tags).Error
	if err != nil {
		return nil, err
	}
	return tags, err
}

func ListVoucher(params *Voucher) []*Voucher {
	var tags []*Voucher
	err := global.DB.Where(&params).Find(&tags).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return tags
}

func GetVoucherById(id uint64) *Voucher {
	var res Voucher
	err := global.DB.Where("id = ?", id).First(&res).Error
	if err != nil {
		return nil
	}
	return &res
}

func CountVoucher() int {
	var count int64
	global.DB.Model(&Voucher{}).Count(&count)
	return int(count)
}
