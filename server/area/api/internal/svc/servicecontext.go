package svc

import (
	"catering/area/api/internal/config"
	"catering/area/rpc/commonhandler"

	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	CommonRpc commonhandler.CommonHandler
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		CommonRpc: commonhandler.NewCommonHandler(zrpc.MustNewClient(c.Common)),
	}
}
