package svc

import (
	snow "catering/area/pkg/snowflake"
	"catering/user/rpc/internal/config"
	"catering/user/rpc/repository"
	"sync"
	"time"

	"github.com/jmoiron/sqlx"
)

type ServiceContext struct {
	Config                config.Config
	UserRepository        repository.UserRepositoryImpl
	UserAddressRepository repository.UserAddressRepositoryImpl
	Snowflake             snow.Snowflake
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := NewMysql(c)
	return &ServiceContext{
		Config: c,
		UserRepository: repository.UserRepositoryImpl{
			Conn: conn,
		},
		UserAddressRepository: repository.UserAddressRepositoryImpl{
			Conn: conn,
		},
	}
}

func NewMysql(c config.Config) *sqlx.DB {
	db := sqlx.MustConnect("mysql", c.Mysql.DataSource)
	db.SetMaxIdleConns(c.Mysql.MaxIdle)
	db.SetMaxOpenConns(c.Mysql.MaxOpen)
	db.SetConnMaxIdleTime(time.Duration(c.Mysql.MaxIdleTime))
	return db
}

func NewSnowflake(c config.Config) *snow.Snowflake {
	var snow *snow.Snowflake
	snow.Datacenterid = c.Snowflake.Datacenterid
	snow.Mutex = sync.Mutex{}
	snow.Sequence = c.Snowflake.Sequence
	snow.Timestamp = c.Snowflake.Timestamp
	snow.Workerid = c.Snowflake.Workerid
	return snow
}
