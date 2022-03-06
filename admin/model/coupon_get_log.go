package model

type CouponGetLog struct {
	Id           uint64  `gorm:"primary_key" json:"id"`
	Uid          int     `json:"uid"`
	CouponId     uint64  `json:"coupon_id"`
	GetTime      Time    `json:"get_time"`
	UseBeginTime Time    `json:"use_begin_time"`
	UseEndTime   Time    `json:"use_end_time"`
	UseStatus    uint32  `json:"use_status"`
	GetType      uint32  `json:"get_type"`
	Coupon       *Coupon `gorm:"foreignKey:CouponId" json:"coupon"`
}

func (CouponGetLog) TableName() string {
	return "user_coupon"
}
