package svc

import (
	"catering/coupon/api/internal/config"
	handler "catering/coupon/rpc/couponhandler"

	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	ConponRpc handler.CouponHandler
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		ConponRpc: handler.NewCouponHandler(zrpc.MustNewClient(c.Coupon)),
	}
}
