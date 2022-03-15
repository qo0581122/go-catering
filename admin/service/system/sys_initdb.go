package system

import (
	"database/sql"
	"fmt"

	"catering/global"
	"catering/model/example"
	"catering/model/system"
	"catering/model/system/request"

	adapter "github.com/casbin/gorm-adapter/v3"
)

type InitDBService struct{}

// InitDB 创建数据库并初始化 总入口
// Author [piexlmax](https://github.com/piexlmax)
// Author [SliverHorn](https://github.com/SliverHorn)
// Author [songzhibin97](https://github.com/songzhibin97)
func (initDBService *InitDBService) InitDB(conf request.InitDB) error {
	switch conf.DBType {
	case "mysql":
		return initDBService.initMsqlDB(conf)
	default:
		return initDBService.initMsqlDB(conf)
	}
}

// initTables 初始化表
// Author [SliverHorn](https://github.com/SliverHorn)
func (initDBService *InitDBService) initTables() error {
	return global.DB.AutoMigrate(
		system.SysApi{},
		system.SysUser{},
		system.SysBaseMenu{},
		system.SysAuthority{},
		system.SysOperationRecord{},
		system.SysBaseMenuParameter{},

		adapter.CasbinRule{},

		example.ExaFile{},
		example.ExaCustomer{},
		example.ExaFileChunk{},
		example.ExaFileUploadAndDownload{},
	)
}

// createDatabase 创建数据库(mysql)
// Author [SliverHorn](https://github.com/SliverHorn)
// Author: [songzhibin97](https://github.com/songzhibin97)

func (initDBService *InitDBService) createDatabase(dsn string, driver string, createSql string) error {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(db)
	if err = db.Ping(); err != nil {
		return err
	}
	_, err = db.Exec(createSql)
	return err
}
