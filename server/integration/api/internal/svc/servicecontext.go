package svc

import (
	"catering/integration/api/internal/config"
	handler "catering/integration/rpc/integrationhandler"

	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	IntegrationRpc handler.IntegrationHandler
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		IntegrationRpc: handler.NewIntegrationHandler(zrpc.MustNewClient(c.Integration)),
	}
}
