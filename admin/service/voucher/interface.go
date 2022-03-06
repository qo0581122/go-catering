package voucher

import (
	"catering/model"
	"catering/model/common/response"
)

type voucherGetLogService interface {
	ListPage(pageNum, pageSize int, params *model.VoucherGetLog) *response.ApiResponse
	GetById(id uint64) *model.VoucherGetLog
	List(params *model.VoucherGetLog) []*model.VoucherGetLog
	Count() int
}

type voucherService interface {
	ListPage(pageNum, pageSize int, params *model.Voucher) *response.ApiResponse
	GetById(id uint64) *model.Voucher
	List(params *model.Voucher) []*model.Voucher
	Count() int
	Add(params *model.Voucher) error
	Update(params *model.Voucher) error
	Delete(id uint64) error
}
