package voucher

import (
	"catering/global"
	"catering/model"
	"catering/model/common/response"
	"errors"
)

var VoucherService voucherService = NewVoucherService()

func NewVoucherService() voucherService {
	return voucherServiceImpl{}
}

type voucherServiceImpl struct {
}

func (impl voucherServiceImpl) Add(params *model.Voucher) error {
	return global.DB.Create(&params).Error
}

func (impl voucherServiceImpl) Delete(id uint64) error {
	query := model.VoucherGetLog{
		VoucherId: id,
	}
	var log model.VoucherGetLog
	if err := global.DB.Where(&query).Find(&log).Error; err != nil {
		return err
	}
	if log.Id > 0 {
		return errors.New("存在关联，无法删除")
	}
	return global.DB.Delete(&model.Voucher{}, id).Error
}

func (impl voucherServiceImpl) Update(params *model.Voucher) error {
	return global.DB.Model(&model.Voucher{}).Where("id = ?", params.ID).Updates(&params).Error
}

func (impl voucherServiceImpl) GetOne(params *model.Voucher) *model.Voucher {
	var res *model.Voucher
	err := global.DB.Where(&params).First(&res).Error
	if err != nil {
		return nil
	}
	return res
}
func (impl voucherServiceImpl) List(params *model.Voucher) []*model.Voucher {
	var vouchers []*model.Voucher
	err := global.DB.Where(&params).Find(&vouchers).Error
	if err != nil {
		return nil
	}
	return vouchers
}

func (impl voucherServiceImpl) Count() int {
	var count int64
	global.DB.Model(&model.Voucher{}).Count(&count)
	return int(count)
}

func (impl voucherServiceImpl) ListPage(pageNum, pageSize int, params *model.Voucher) *response.ApiResponse {
	var vouchers []*model.Voucher
	err := global.DB.Where(&params).Find(&vouchers).Error
	if err != nil {
		return nil
	}
	var count int64
	global.DB.Model(&model.Voucher{}).Where(&params).Count(&count)
	res := &response.ApiResponse{
		List:  vouchers,
		Total: int(count),
	}
	return res
}
