package svc

import (
	"catering/integration/rpc/internal/config"
	"catering/integration/rpc/repository"
	"catering/pkg/mysql"

	"github.com/tal-tech/go-zero/core/stores/cache"
)

type ServiceContext struct {
	Config                config.Config
	Cache                 *cache.Cache
	IntegrationRepository repository.IntegrationRepository
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := mysql.NewMysql(mysql.MysqlConfig(&c))
	return &ServiceContext{
		Config: c,
		IntegrationRepository: &repository.IntegrationRepositoryImpl{
			Conn: conn,
		},
	}
}
