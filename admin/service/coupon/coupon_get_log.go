package coupon

import (
	"catering/model"
	"catering/model/common/response"
	"fmt"
)

type CouponGetLogDetail struct {
	Log    *model.CouponGetLog `json:"log"`
	Coupon *model.Coupon       `json:"coupon"`
}

var CouponGetLogService couponGetLogService = NewCouponGetLogService()

func NewCouponGetLogService() couponGetLogService {
	return couponGetLogServiceImpl{}
}

type couponGetLogServiceImpl struct {
}

func (impl couponGetLogServiceImpl) GetById(id uint64) *model.CouponGetLog {
	return model.GetCouponGetLogById(id)
}
func (impl couponGetLogServiceImpl) List(params *model.CouponGetLog) []*model.CouponGetLog {
	return model.ListCouponGetLog(params)
}

func (impl couponGetLogServiceImpl) Count() int {
	return model.CountUserAddress()
}

func (impl couponGetLogServiceImpl) ListPage(pageNum, pageSize int, params *model.CouponGetLog) *response.ApiResponse {
	logs, err := model.ListCouponGetLogPage(pageNum, pageSize, params)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var result []*CouponGetLogDetail
	for _, item := range logs {
		couponId := item.CouponId
		coupon := model.GetCouponById(couponId)
		result = append(result, &CouponGetLogDetail{
			Log:    item,
			Coupon: coupon,
		})
	}
	total := impl.Count()
	res := &response.ApiResponse{List: result, Total: total}
	return res
}
