package model

import (
	"catering/global"
	"fmt"
)

type VoucherGetLog struct {
	Id           uint64 `gorm:"primary_key" json:"id"`
	Uid          int    `json:"uid"`
	VoucherId    uint64 `json:"voucher_id"`
	GetTime      Time   `json:"get_time"`
	UseBeginTime Time   `json:"use_begin_time"`
	UseEndTime   Time   `json:"use_end_time"`
	UseStatus    uint32 `json:"use_status"`
	GetType      uint32 `json:"get_type"`
}

func (VoucherGetLog) TableName() string {
	return "user_voucher"
}

func AddVoucherGetLog(params *VoucherGetLog) error {
	return global.DB.Create(&params).Error
}

func UpdateVoucherGetLog(params *VoucherGetLog) error {
	return global.DB.Model(&VoucherGetLog{}).Where("id = ?", params.Id).Updates(&params).Error
}

func DeleteVoucherGetLog(id uint64) error {
	return global.DB.Delete(&VoucherGetLog{}, id).Error
}

func ListVoucherGetLogPage(pageNum, pageSize int, params *VoucherGetLog) ([]*VoucherGetLog, error) {
	var tags []*VoucherGetLog
	err := global.DB.Where(&params).Scopes(Paginate(pageNum, pageSize)).Find(&tags).Error
	if err != nil {
		return nil, err
	}
	return tags, err
}

func ListVoucherGetLog(params *VoucherGetLog) []*VoucherGetLog {
	var tags []*VoucherGetLog
	err := global.DB.Where(&params).Find(&tags).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return tags
}

func GetVoucherGetLogById(id uint64) *VoucherGetLog {
	var res VoucherGetLog
	err := global.DB.Where("id = ?", id).First(&res).Error
	if err != nil {
		return nil
	}
	return &res
}

func CountVoucherGetLog() int {
	var count int64
	global.DB.Model(&VoucherGetLog{}).Count(&count)
	return int(count)
}
