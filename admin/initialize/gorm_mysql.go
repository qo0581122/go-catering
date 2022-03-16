package initialize

import (
	"catering/global"
	"catering/initialize/internal"
	"database/sql"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GormMysql() *gorm.DB {
	cfg := global.Config.Mysql
	if err := cfg.Check(); err != nil {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       cfg.Dsn(), // DSN data source name
		DefaultStringSize:         256,       // string 类型字段的默认长度
		DisableDatetimePrecision:  true,      // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,      // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,      // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,     // 根据当前 MySQL 版本自动配置
	}
	db, err := gorm.Open(mysql.New(mysqlConfig), internal.Gorm.Config())
	if err != nil {
		return nil
	}
	sqlDB, _ := db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	Options(sqlDB, WithMaxIdelConns(cfg.MaxIdleConns), WithMaxOpenConns(cfg.MaxOpenConns), WithMaxLifeTime(cfg.MaxLifeTime))
	return db
}

type Option func(m *sql.DB)

func WithMaxIdelConns(idle int) Option {
	return func(m *sql.DB) {
		if idle == 0 {
			return
		}
		m.SetMaxIdleConns(idle)
	}
}

func WithMaxOpenConns(open int) Option {
	return func(m *sql.DB) {
		if open == 0 {
			return
		}
		m.SetMaxOpenConns(open)
	}
}

func WithMaxLifeTime(t int) Option {
	return func(m *sql.DB) {
		if t == 0 {
			return
		}
		m.SetConnMaxLifetime(time.Duration(t) * time.Second)
	}
}
func Options(m *sql.DB, opts ...Option) {
	for _, opt := range opts {
		opt(m)
	}
}
