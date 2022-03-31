package coupon

import (
	"catering/global"
	"catering/model"
	"catering/model/common/response"
	"errors"
	"fmt"
)

var CouponService couponService = NewCouponService()

func NewCouponService() couponService {
	return couponServiceImpl{}
}

type couponServiceImpl struct {
}

func (impl couponServiceImpl) Add(params *model.Coupon) error {
	return global.DB.Create(&params).Error
}
func (impl couponServiceImpl) Delete(id uint64) error {
	query := model.CouponGetLog{
		CouponId: id,
	}
	var log model.CouponGetLog
	if err := global.DB.Where(&query).Find(&log).Error; err != nil {
		return err
	}
	if log.Id > 0 {
		return errors.New("存在关联，无法删除")
	}
	return global.DB.Delete(&model.Coupon{}, id).Error
}
func (impl couponServiceImpl) Update(params *model.Coupon) error {
	return global.DB.Model(&model.Coupon{}).Where("id = ?", params.ID).Updates(&params).Error
}

func (impl couponServiceImpl) GetOne(params *model.Coupon) *model.Coupon {
	var res model.Coupon
	err := global.DB.Where(&params).First(&res).Error
	if err != nil {
		return nil
	}
	return &res
}

func (impl couponServiceImpl) List(params *model.Coupon) []*model.Coupon {
	var coupons []*model.Coupon
	err := global.DB.Where(&params).Find(&coupons).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return coupons
}

func (impl couponServiceImpl) Count() int {
	var count int64
	global.DB.Model(&model.Coupon{}).Count(&count)
	return int(count)
}

func (impl couponServiceImpl) ListPage(pageNum, pageSize int, params *model.Coupon) *response.ApiResponse {
	var coupons []*model.Coupon
	if params.Name != "" {
		global.DB = global.DB.Where("coupon_name LIKE ?", "%"+params.Name+"%")
		params.Name = ""
	}
	err := global.DB.Preload("Product").Where(&params).Scopes(model.Paginate(pageNum, pageSize)).Find(&coupons).Error
	if err != nil {
		return nil
	}

	var count int64
	global.DB.Model(&model.Coupon{}).Count(&count)
	return &response.ApiResponse{
		List:  coupons,
		Total: int(count),
	}
}
