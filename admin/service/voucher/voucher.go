package voucher

import (
	"catering/model"
	"catering/model/common/response"
	"fmt"
)

var VoucherService voucherService = NewVoucherService()

func NewVoucherService() voucherService {
	return voucherServiceImpl{}
}

type voucherServiceImpl struct {
}

func (impl voucherServiceImpl) Add(params *model.Voucher) error {
	return model.AddVoucher(params)
}
func (impl voucherServiceImpl) Delete(id uint64) error {
	return model.DeleteVoucher(id)
}
func (impl voucherServiceImpl) Update(params *model.Voucher) error {
	return model.UpdateVoucher(params)
}

func (impl voucherServiceImpl) GetById(id uint64) *model.Voucher {
	return model.GetVoucherById(id)
}
func (impl voucherServiceImpl) List(params *model.Voucher) []*model.Voucher {
	return model.ListVoucher(params)
}

func (impl voucherServiceImpl) Count() int {
	return model.CountUserAddress()
}

func (impl voucherServiceImpl) ListPage(pageNum, pageSize int, params *model.Voucher) *response.ApiResponse {
	vouchers, err := model.ListVoucherPage(pageNum, pageSize, params)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	count := model.CountVoucher()
	res := &response.ApiResponse{
		List:  vouchers,
		Total: count,
	}
	return res
}
