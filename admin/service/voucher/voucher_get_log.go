package voucher

import (
	"catering/model"
	"catering/model/common/response"
	"fmt"
)

var VoucherGetLogService voucherGetLogService = NewVoucherGetLogService()

type VoucherGetLogDetail struct {
	Log     *model.VoucherGetLog `json:"log"`
	Voucher *model.Voucher       `json:"voucher"`
}

func NewVoucherGetLogService() voucherGetLogService {
	return voucherGetLogServiceImpl{}
}

type voucherGetLogServiceImpl struct {
}

func (impl voucherGetLogServiceImpl) GetById(id uint64) *model.VoucherGetLog {
	return model.GetVoucherGetLogById(id)
}
func (impl voucherGetLogServiceImpl) List(params *model.VoucherGetLog) []*model.VoucherGetLog {
	return model.ListVoucherGetLog(params)
}

func (impl voucherGetLogServiceImpl) Count() int {
	return model.CountUserAddress()
}

func (impl voucherGetLogServiceImpl) ListPage(pageNum, pageSize int, params *model.VoucherGetLog) *response.ApiResponse {
	logs, err := model.ListVoucherGetLogPage(pageNum, pageSize, params)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var result []*VoucherGetLogDetail
	for _, item := range logs {
		voucherId := item.VoucherId
		voucher := model.GetVoucherById(voucherId)
		result = append(result, &VoucherGetLogDetail{
			Log:     item,
			Voucher: voucher,
		})
	}
	total := impl.Count()
	res := &response.ApiResponse{List: result, Total: total}
	return res
}
