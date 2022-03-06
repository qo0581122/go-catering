package model

import (
	"catering/global"
	"fmt"
)

type CouponDetail struct {
	Coupon  *Coupon  `json:"coupon"`
	Product *Product `json:"product"`
}

type Coupon struct {
	Model
	Name          string `gorm:"column:coupon_name" json:"coupon_name"`
	Price         uint64 `json:"price"`
	ProductId     uint64 `json:"product_id"`
	LeastUsePrice uint64 `json:"least_use_price"`
	PicUrl        string `json:"pic_url"`
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
	Description   string `json:"description"`
}

func (Coupon) TableName() string {
	return "coupon"
}

func AddCoupon(params *Coupon) error {
	return global.DB.Create(&params).Error
}

func UpdateCoupon(params *Coupon) error {
	return global.DB.Model(&Coupon{}).Where("id = ?", params.Id).Updates(&params).Error
}

func DeleteCoupon(id uint64) error {
	return global.DB.Delete(&Coupon{}, id).Error
}

func ListCouponPage(pageNum, pageSize int, params *Coupon) ([]*Coupon, error) {
	var coupons []*Coupon
	if params.Name != "" {
		global.DB = global.DB.Where("coupon_name LIKE ?", "%"+params.Name+"%")
		params.Name = ""
	}
	err := global.DB.Where(&params).Scopes(Paginate(pageNum, pageSize)).Find(&coupons).Error
	if err != nil {
		return nil, err
	}
	return coupons, err
}

func ListCoupon(params *Coupon) []*Coupon {
	var coupons []*Coupon
	err := global.DB.Where(&params).Find(&coupons).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return coupons
}

func CountCoupon() int {
	var count int64
	global.DB.Model(&Coupon{}).Count(&count)
	return int(count)
}

func GetCouponById(id uint64) *Coupon {
	var res Coupon
	err := global.DB.Where("id = ?", id).First(&res).Error
	if err != nil {
		return nil
	}
	return &res
}
