package mysql

import (
	"github.com/jmoiron/sqlx"
)

func NewMysql(c MysqlConfig) *sqlx.DB {
	db := sqlx.MustConnect("mysql", c.GetDataSource())
	db.SetMaxIdleConns(c.GetMaxIdle())
	db.SetMaxOpenConns(c.GetMaxOpen())
	return db
}

type MysqlConfig interface {
	GetDataSource() string
	GetMaxIdle() int
	GetMaxOpen() int
}
