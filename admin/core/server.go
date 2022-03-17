package core

import (
	"fmt"
	"time"

	"catering/global"
	"catering/initialize"

	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {

	Router := initialize.Routers()

	Router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.Config.System.Addr)
	s := initServer(address, Router)
	// 保证文本顺序输出
	time.Sleep(10 * time.Microsecond)
	global.Log.Info("server run success on ", zap.String("address", address))

	fmt.Printf(`
	欢迎使用 catering
`, address)
	global.Log.Error(s.ListenAndServe().Error())
}
