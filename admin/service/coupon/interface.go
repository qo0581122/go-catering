package coupon

import (
	"catering/model"
	"catering/model/common/response"
)

type couponGetLogService interface {
	ListPage(pageNum, pageSize int, params *model.CouponGetLog) *response.ApiResponse
	GetOne(params *model.CouponGetLog) *model.CouponGetLog
	List(params *model.CouponGetLog) []*model.CouponGetLog
	Count() int
}

type couponService interface {
	ListPage(pageNum, pageSize int, params *model.Coupon) *response.ApiResponse
	GetOne(params *model.Coupon) *model.Coupon
	List(params *model.Coupon) []*model.Coupon
	Count() int
	Add(params *model.Coupon) error
	Update(params *model.Coupon) error
	Delete(id uint64) error
}
