package coupon

import (
	"catering/model"
	"catering/model/common/response"
	"fmt"
)

var CouponService couponService = NewCouponService()

func NewCouponService() couponService {
	return couponServiceImpl{}
}

type couponServiceImpl struct {
}

func (impl couponServiceImpl) Add(params *model.Coupon) error {
	return model.AddCoupon(params)
}
func (impl couponServiceImpl) Delete(id uint64) error {
	return model.DeleteCoupon(id)
}
func (impl couponServiceImpl) Update(params *model.Coupon) error {
	return model.UpdateCoupon(params)
}

func (impl couponServiceImpl) GetById(id uint64) *model.Coupon {
	return model.GetCouponById(id)
}
func (impl couponServiceImpl) List(params *model.Coupon) []*model.Coupon {
	return model.ListCoupon(params)
}

func (impl couponServiceImpl) Count() int {
	return model.CountUserAddress()
}

func (impl couponServiceImpl) ListPage(pageNum, pageSize int, params *model.Coupon) *response.ApiResponse {
	coupons, err := model.ListCouponPage(pageNum, pageSize, params)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var detail []*model.CouponDetail
	for _, item := range coupons {
		productId := item.ProductId
		product := model.GetProductById(productId)
		detail = append(detail, &model.CouponDetail{
			Coupon:  item,
			Product: product,
		})
	}
	count := model.CountCoupon()
	res := &response.ApiResponse{
		List:  detail,
		Total: count,
	}
	return res
}

func (impl couponServiceImpl) ListByProductId(pid uint64) []*model.CouponDetail {
	var result []*model.CouponDetail
	coupons := model.ListCoupon(&model.Coupon{
		ProductId: pid,
	})
	if len(result) <= 0 {
		return result
	}
	for _, item := range coupons {
		productId := item.ProductId
		product := model.GetProductById(productId)
		result = append(result, &model.CouponDetail{
			Coupon:  item,
			Product: product,
		})
	}
	return result
}
