package svc

import (
	"catering/user/api/internal/config"
	"catering/user/rpc/userhandler"

	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	UserService userhandler.UserHandler
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		UserService: userhandler.NewUserHandler(zrpc.MustNewClient(c.User)),
	}
}
