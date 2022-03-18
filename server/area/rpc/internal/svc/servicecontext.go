package svc

import (
	"catering/area/rpc/internal/config"
	"catering/area/rpc/repository"

	"github.com/jmoiron/sqlx"
)

type ServiceContext struct {
	Config             config.Config
	ProvinceRepository repository.ProvinceRepositoryImpl
	CityRepository     repository.CityRepositoryImpl
	DistinctRepository repository.DistinctRepositoryImpl
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := NewMysql(c)
	return &ServiceContext{
		Config: c,
		ProvinceRepository: repository.ProvinceRepositoryImpl{
			Conn: conn,
		},
		CityRepository: repository.CityRepositoryImpl{
			Conn: conn,
		},
		DistinctRepository: repository.DistinctRepositoryImpl{
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
