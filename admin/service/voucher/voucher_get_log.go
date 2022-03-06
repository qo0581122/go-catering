package voucher

import (
	"catering/global"
	"catering/model"
	"catering/model/common/response"
	"fmt"
)

var VoucherGetLogService voucherGetLogService = NewVoucherGetLogService()

func NewVoucherGetLogService() voucherGetLogService {
	return voucherGetLogServiceImpl{}
}

type voucherGetLogServiceImpl struct {
}

func (impl voucherGetLogServiceImpl) GetOne(params *model.VoucherGetLog) *model.VoucherGetLog {
	var res model.VoucherGetLog
	err := global.DB.Where(&params).First(&res).Error
	if err != nil {
		return nil
	}
	return &res
}
func (impl voucherGetLogServiceImpl) List(params *model.VoucherGetLog) []*model.VoucherGetLog {
	var logs []*model.VoucherGetLog
	err := global.DB.Where(&params).Find(&logs).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return logs
}

func (impl voucherGetLogServiceImpl) Count() int {
	var count int64
	global.DB.Model(&model.VoucherGetLog{}).Count(&count)
	return int(count)
}

func (impl voucherGetLogServiceImpl) ListPage(pageNum, pageSize int, params *model.VoucherGetLog) *response.ApiResponse {
	var logs []*model.VoucherGetLog
	err := global.DB.Preload("Voucher").Where(&params).Scopes(model.Paginate(pageNum, pageSize)).Find(&logs).Error
	if err != nil {
		return nil
	}
	var total int64
	global.DB.Model(&model.VoucherGetLog{}).Where(&params).Count(&total)
	res := &response.ApiResponse{List: logs, Total: int(total)}
	return res
}
