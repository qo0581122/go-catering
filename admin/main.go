package main

import (
	"catering/core"
	"catering/global"
	"catering/initialize"
	"fmt"
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
	err := InitOptions(initialize.InitConfig, initialize.InitZap, initialize.InitGorm, initialize.InitSession)
	if err != nil {
		fmt.Println("server run fail:", err.Error())
		return
	}
	if global.DB != nil {
		db, _ := global.DB.DB()
		defer db.Close()
	}
	core.RunWindowsServer()
}

type InitFuc func() error

func InitOptions(fuc ...InitFuc) error {
	for _, f := range fuc {
		err := f()
		if err != nil {
			return err
		}
	}
	return nil
}
