package initialize

import (
	"catering/global"
	"catering/initialize/internal"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// GormMysql 初始化Mysql数据库
// Author [piexlmax](https://github.com/piexlmax)
// Author [SliverHorn](https://github.com/SliverHorn)
// func GormMysql() *gorm.DB {
// 	m := global.Config.Mysql
// 	if m.Dbname == "" {
// 		return nil
// 	}
// 	mysqlConfig := mysql.Config{
// 		DSN:                       m.Dsn(), // DSN data source name
// 		DefaultStringSize:         191,     // string 类型字段的默认长度
// 		SkipInitializeWithVersion: false,   // 根据版本自动配置
// 	}
// 	if db, err := gorm.Open(mysql.New(mysqlConfig), internal.Gorm.Config()); err != nil {
// 		return nil
// 	} else {
// 		sqlDB, _ := db.DB()
// 		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
// 		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
// 		return db
// 	}
// }

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
	if cfg.MaxIdleConns != 0 {
		sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	}
	if cfg.MaxOpenConns != 0 {
		sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	}
	if cfg.MaxLifeTime != 0 {
		sqlDB.SetConnMaxLifetime(time.Duration(cfg.MaxLifeTime) * time.Second)
	}
	return db
}
