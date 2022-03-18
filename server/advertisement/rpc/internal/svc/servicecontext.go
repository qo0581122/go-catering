package svc

import (
	"catering/advertisement/rpc/internal/config"
	"catering/advertisement/rpc/repository"
	ca "catering/area/pkg/cache"
	"catering/product/producthandler"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config                     config.Config
	Cache                      *cache.Cache
	AdBannerRepository         repository.AdBannerRepositoryImpl
	AdHotActivitieRepository   repository.AdHotActivitieRepositoryImpl
	AdRecommendDailyRepository repository.AdRecommendDailyRepositoryImpl
	ProductService             producthandler.ProductHandler
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := NewMysql(c)
	return &ServiceContext{
		Config: c,
		Cache:  ca.NewCache(c.Cache),
		AdBannerRepository: repository.AdBannerRepositoryImpl{
			Conn: conn,
		},
		AdHotActivitieRepository: repository.AdHotActivitieRepositoryImpl{
			Conn: conn,
		},
		AdRecommendDailyRepository: repository.AdRecommendDailyRepositoryImpl{
			Conn: conn,
		},
		ProductService: producthandler.NewProductHandler(zrpc.MustNewClient(c.Product)),
	}
}

func NewMysql(c config.Config) *sqlx.DB {
	db := sqlx.MustConnect("mysql", c.Mysql.DataSource)
	db.SetMaxIdleConns(c.Mysql.MaxIdle)
	db.SetMaxOpenConns(c.Mysql.MaxOpen)
	db.SetConnMaxIdleTime(time.Duration(c.Mysql.MaxIdle))
	return db
}
