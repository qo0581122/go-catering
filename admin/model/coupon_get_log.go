package model

import (
	"catering/global"
	"fmt"
)

type CouponGetLog struct {
	Id           uint64 `gorm:"primary_key" json:"id"`
	Uid          int    `json:"uid"`
	CouponId     uint64 `json:"coupon_id"`
	GetTime      Time   `json:"get_time"`
	UseBeginTime Time   `json:"use_begin_time"`
	UseEndTime   Time   `json:"use_end_time"`
	UseStatus    uint32 `json:"use_status"`
	GetType      uint32 `json:"get_type"`
}

func (CouponGetLog) TableName() string {
	return "user_coupon"
}

func ListCouponGetLogPage(pageNum, pageSize int, params *CouponGetLog) ([]*CouponGetLog, error) {
	var coupons []*CouponGetLog
	err := global.DB.Where(&params).Scopes(Paginate(pageNum, pageSize)).Find(&coupons).Error
	if err != nil {
		return nil, err
	}
	return coupons, err
}

func ListCouponGetLog(params *CouponGetLog) []*CouponGetLog {
	var coupons []*CouponGetLog
	err := global.DB.Where(&params).Find(&coupons).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return coupons
}

func CountCouponGetLog() int {
	var count int64
	global.DB.Model(&CouponGetLog{}).Count(&count)
	return int(count)
}

func GetCouponGetLogById(id uint64) *CouponGetLog {
	var res CouponGetLog
	err := global.DB.Where("id = ?", id).First(&res).Error
	if err != nil {
		return nil
	}
	return &res
}
