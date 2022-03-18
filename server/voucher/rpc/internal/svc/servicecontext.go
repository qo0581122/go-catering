package svc

import (
	"catering/voucher/rpc/internal/config"
	"catering/voucher/rpc/repository"

	"github.com/jmoiron/sqlx"
	"github.com/tal-tech/go-zero/core/stores/cache"
)

type ServiceContext struct {
	Config            config.Config
	Cache             *cache.Cache
	VoucherRepository repository.VoucherRepository
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := NewMysql(c)
	return &ServiceContext{
		Config: c,
		VoucherRepository: &repository.VoucherRepositoryImpl{
			Conn: conn,
		},
	}
}

func NewMysql(c config.Config) *sqlx.DB {
	db := sqlx.MustConnect("mysql", c.Mysql.DataSource)
	db.SetMaxIdleConns(c.Mysql.MaxIdle)
	db.SetMaxOpenConns(c.Mysql.MaxOpen)
	return db
}
