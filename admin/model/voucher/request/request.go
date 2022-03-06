package request

import "catering/model/common/request"

type VoucherQueryParams struct {
	request.PageParams
	VoucherName string `uri:"voucher_name"  form:"voucher_name" json:"voucher_name"`
	Status      uint32 `uri:"status" json:"status" form:"status"`
}

type VoucherAddForm struct {
	VoucherName   string `json:"voucher_name" form:"voucher_name" valid:"Required"`
	Price         uint64 `json:"price" form:"price" valid:"Required"`
	LeastUsePrice uint64 `json:"least_use_price" form:"least_use_price" `
	PicUrl        string `json:"pic_url" form:"pic_url" `
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

type VoucherUpdateForm struct {
	Id            uint64 `form:"id" json:"id" valid:"Required"`
	VoucherName   string `json:"voucher_name" form:"voucher_name" valid:"Required"`
	Price         uint64 `json:"price" form:"price" valid:"Required"`
	LeastUsePrice uint64 `json:"least_use_price" form:"least_use_price" `
	PicUrl        string `json:"pic_url" form:"pic_url" `
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

type VoucherGetLogQueryParams struct {
	request.PageParams
	Uid       int    `uri:"uid" json:"uid" form:"uid"`
	VoucherId uint64 `uri:"voucher_id" json:"voucher_id" form:"voucher_id"`
	UseStatus uint32 `uri:"use_status" json:"use_status" form:"use_status"`
	GetType   uint32 `uri:"get_type" json:"get_type" form:"get_type"`
}
