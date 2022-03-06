package initialize

import (
	"catering/global"

	"gorm.io/gorm"
)

// Gorm 初始化数据库并产生数据库全局变量
// Author SliverHorn
func InitGorm() {
	var db *gorm.DB
	switch global.Config.System.DbType {
	case "mysql":
		db = GormMysql()
	default:
		db = GormMysql()
	}

	global.DB = db
}
