package main

import (
	"catering/core"
	"catering/global"
	"catering/initialize"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func main() {
	initialize.InitConfig()
	initialize.InitZap()
	initialize.InitGorm()
	initialize.InitSession()
	// 初始化redis服务
	// initialize.Redis()
	if global.DB != nil {
		db, _ := global.DB.DB()
		defer db.Close()
	}
	core.RunWindowsServer()
}
