package coupon

import (
	"catering/global"
	"catering/model"
	"catering/model/common/response"
	"fmt"
)

var CouponGetLogService couponGetLogService = NewCouponGetLogService()

func NewCouponGetLogService() couponGetLogService {
	return couponGetLogServiceImpl{}
}

type couponGetLogServiceImpl struct {
}

func (impl couponGetLogServiceImpl) GetOne(params *model.CouponGetLog) *model.CouponGetLog {
	var res *model.CouponGetLog
	err := global.DB.Where(&params).First(&res).Error
	if err != nil {
		return nil
	}
	return res
}
func (impl couponGetLogServiceImpl) List(params *model.CouponGetLog) []*model.CouponGetLog {
	var coupons []*model.CouponGetLog
	err := global.DB.Where(&params).Find(&coupons).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return coupons
}

func (impl couponGetLogServiceImpl) Count() int {
	var count int64
	global.DB.Model(&model.CouponGetLog{}).Count(&count)
	return int(count)
}

func (impl couponGetLogServiceImpl) ListPage(pageNum, pageSize int, params *model.CouponGetLog) *response.ApiResponse {
	var logs []*model.CouponGetLog
	err := global.DB.Preload("Coupon").Where(&params).Scopes(model.Paginate(pageNum, pageSize)).Find(&logs).Error
	if err != nil {
		return nil
	}
	var total int64
	global.DB.Model(&model.CouponGetLog{}).Count(&total)
	res := &response.ApiResponse{List: logs, Total: int(total)}
	return res
}
