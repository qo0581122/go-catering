package system

import (
	"catering/config"
	"catering/global"
	"catering/model/system"
	"catering/pkg"

	"go.uber.org/zap"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetSystemConfig
//@description: 读取配置文件
//@return: err error, conf config.Server

type SystemConfigService struct{}

func (systemConfigService *SystemConfigService) GetSystemConfig() (err error, conf config.Config) {
	return nil, global.Config
}

// @description   set system config,
//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetSystemConfig
//@description: 设置配置文件
//@param: system model.System
//@return: err error

func (systemConfigService *SystemConfigService) SetSystemConfig(system system.System) (err error) {
	// cs := utils.StructToMap(system.Config)
	// for k, v := range cs {
	// 	global.GVA_VP.Set(k, v)
	// }
	// err = global.GVA_VP.WriteConfig()
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: GetServerInfo
//@description: 获取服务器信息
//@return: server *utils.Server, err error

func (systemConfigService *SystemConfigService) GetServerInfo() (server *pkg.Server, err error) {
	var s pkg.Server
	s.Os = pkg.InitOS()
	if s.Cpu, err = pkg.InitCPU(); err != nil {
		global.Log.Error("func utils.InitCPU() Failed", zap.String("err", err.Error()))
		return &s, err
	}
	if s.Rrm, err = pkg.InitRAM(); err != nil {
		global.Log.Error("func utils.InitRAM() Failed", zap.String("err", err.Error()))
		return &s, err
	}
	if s.Disk, err = pkg.InitDisk(); err != nil {
		global.Log.Error("func utils.InitDisk() Failed", zap.String("err", err.Error()))
		return &s, err
	}

	return &s, nil
}
