package svc

import (
	ca "catering/area/pkg/cache"
	"catering/product/internal/config"
	"catering/product/repository"

	"github.com/jmoiron/sqlx"
	"github.com/tal-tech/go-zero/core/stores/cache"
)

type ServiceContext struct {
	Config                     config.Config
	ProductAttributeRepository repository.ProductAttributeRepositoryImpl
	ProductBaseRepository      repository.ProductBaseRepositoryImpl
	ProductBatchRepository     repository.ProductBatchRepositoryImpl
	ProductCategoryRepository  repository.ProductCategoryRepositoryImpl
	ProductRelationRepository  repository.ProductRelationRepositoryImpl
	Cache                      *cache.Cache
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := NewMysql(c)
	return &ServiceContext{
		Config: c,
		ProductAttributeRepository: repository.ProductAttributeRepositoryImpl{
			Conn: conn,
		},
		ProductBaseRepository: repository.ProductBaseRepositoryImpl{
			Conn: conn,
		},
		ProductBatchRepository: repository.ProductBatchRepositoryImpl{
			Conn: conn,
		},
		ProductCategoryRepository: repository.ProductCategoryRepositoryImpl{
			Conn: conn,
		},
		ProductRelationRepository: repository.ProductRelationRepositoryImpl{
			Conn: conn,
		},
		Cache: ca.NewCache(c.Cache),
	}
}

func NewMysql(c config.Config) *sqlx.DB {
	db := sqlx.MustConnect("mysql", c.Mysql.DataSource)
	db.SetMaxIdleConns(c.Mysql.MaxIdle)
	db.SetMaxOpenConns(c.Mysql.MaxOpen)
	return db
}
