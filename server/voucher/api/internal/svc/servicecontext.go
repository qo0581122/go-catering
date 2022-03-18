package svc

import (
	"catering/voucher/api/internal/config"
	. "catering/voucher/rpc/voucherhandler"

	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	VoucherRpc VoucherHandler
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		VoucherRpc: NewVoucherHandler(zrpc.MustNewClient(c.Voucher)),
	}
}
