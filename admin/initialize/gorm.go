package initialize

import (
	"catering/global"

	"gorm.io/gorm"
)

// Gorm 初始化数据库并产生数据库全局变量
// Author SliverHorn
func InitGorm() error {
	var db *gorm.DB
	var err error
	switch global.Config.System.DbType {
	case "mysql":
		db, err = GormMysql()
	default:
		db, err = GormMysql()
	}
	if err != nil {
		return err
	}
	global.DB = db
	return nil
}
