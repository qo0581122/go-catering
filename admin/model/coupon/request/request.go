package request

import "catering/model/common/request"

type CouponQueryParams struct {
	request.PageParams
	PageSize   int    `uri:"pageSize" json:"pageSize" form:"pageSize" valid:"Required"`
	PageNum    int    `uri:"pageNum" json:"pageNum" form:"pageNum" valid:"Required"`
	ProductId  uint64 `uri:"product_id"  form:"product_id" json:"product_id"`
	CouponName string `uri:"coupon_name"  form:"coupon_name" json:"coupon_name"`
	Status     uint32 `uri:"status" json:"status" form:"status"`
}

type CouponAddForm struct {
	CouponName    string `json:"coupon_name" form:"coupon_name" valid:"Required"`
	Price         uint64 `json:"price" form:"price" valid:"Required"`
	ProductId     uint64 `json:"product_id" form:"product_id" valid:"Required"`
	LeastUsePrice uint64 `json:"least_use_price" form:"least_use_price" `
	TotalCount    uint64 `json:"total_count" form:"total_count" valid:"Required"`
	GetType       uint32 `json:"get_type" form:"get_type" valid:"Required"`
	ValidTimeType uint32 `json:"valid_time_type" form:"valid_time_type"  `
	ValidDay      uint32 `json:"valid_day" form:"valid_day" `
	UseType       uint32 `json:"use_type" form:"use_type" valid:"Required"`
	Status        uint32 `json:"status" form:"status" `
	GetBeginTime  string `json:"get_begin_time" form:"get_begin_time" valid:"Required"`
	GetEndTime    string `json:"get_end_time" form:"get_end_time" valid:"Required"`
	UseBeginTime  string `json:"use_begin_time" form:"use_begin_time" `
	UseEndTime    string `json:"use_end_time" form:"use_end_time" `
	GetCount      uint32 `json:"get_count" form:"get_count" valid:"Required"`
	Description   string `json:"description" form:"description" `
}

type CouponUpdateForm struct {
	Id            uint64 `form:"id" json:"id" valid:"Required"`
	CouponName    string `json:"coupon_name" form:"coupon_name" valid:"Required"`
	Price         uint64 `json:"price" form:"price" valid:"Required"`
	ProductId     uint64 `json:"product_id" form:"product_id" valid:"Required"`
	LeastUsePrice uint64 `json:"least_use_price" form:"least_use_price" `
	TotalCount    uint64 `json:"total_count" form:"total_count" valid:"Required"`
	RemainCount   uint64 `json:"remain_count" form:"remain_count" valid:"Required"`
	GetType       uint32 `json:"get_type" form:"get_type" valid:"Required"`
	ValidTimeType int    `json:"valid_time_type" form:"valid_time_type"  `
	ValidDay      uint32 `json:"valid_day" form:"valid_day" `
	UseType       uint32 `json:"use_type" form:"use_type" valid:"Required"`
	Status        int    `json:"status" form:"status" `
	GetBeginTime  string `json:"get_begin_time" form:"get_begin_time" valid:"Required"`
	GetEndTime    string `json:"get_end_time" form:"get_end_time" valid:"Required"`
	UseBeginTime  string `json:"use_begin_time" form:"use_begin_time" `
	UseEndTime    string `json:"use_end_time" form:"use_end_time" `
	GetCount      uint32 `json:"get_count" form:"get_count" valid:"Required"`
	Description   string `json:"description" form:"description" `
}

type CouponGetLogQueryParams struct {
	request.PageParams
	Uid       int    `uri:"uid" json:"uid" form:"uid"`
	CouponId  uint64 `uri:"coupon_id" json:"coupon_id" form:"coupon_id"`
	UseStatus uint32 `uri:"use_status" json:"use_status" form:"use_status"`
	GetType   uint32 `uri:"get_type" json:"get_type" form:"get_type"`
}
