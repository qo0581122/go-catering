package svc

import (
	"catering/coupon/rpc/internal/config"
	"catering/coupon/rpc/repository"
	"catering/product/producthandler"

	"github.com/jmoiron/sqlx"
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config           config.Config
	Cache            *cache.Cache
	CouponRepository repository.CouponRepository
	ProductService   producthandler.ProductHandler
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := NewMysql(c)
	return &ServiceContext{
		Config: c,
		CouponRepository: &repository.CouponRepositoryImpl{
			Conn: conn,
		},
		ProductService: producthandler.NewProductHandler(zrpc.MustNewClient(c.Product)),
	}
}

func NewMysql(c config.Config) *sqlx.DB {
	db := sqlx.MustConnect("mysql", c.Mysql.DataSource)
	db.SetMaxIdleConns(c.Mysql.MaxIdle)
	db.SetMaxOpenConns(c.Mysql.MaxOpen)
	return db
}
