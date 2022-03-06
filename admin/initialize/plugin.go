package initialize

import (
	"catering/global"

	email "github.com/flipped-aurora/gva-plugins/email" // 在线仓库模式go

	//"catering/plugin/email" // 本地插件仓库地址模式
	"catering/plugin/example_plugin"
	"catering/pkg/plugin"

	"github.com/gin-gonic/gin"
)

func PluginInit(group *gin.RouterGroup, Plugin ...plugin.Plugin) {
	for i := range Plugin {
		PluginGroup := group.Group(Plugin[i].RouterPath())
		Plugin[i].Register(PluginGroup)
	}
}

func InstallPlugin(PublicGroup *gin.RouterGroup, PrivateGroup *gin.RouterGroup) {
	//  添加开放权限的插件 示例
	PluginInit(PublicGroup,
		example_plugin.ExamplePlugin)

	//  添加跟角色挂钩权限的插件 示例 本地示例模式于在线仓库模式注意上方的import 可以自行切换 效果相同
	PluginInit(PrivateGroup, email.CreateEmailPlug(
		global.Config.Email.To,
		global.Config.Email.From,
		global.Config.Email.Host,
		global.Config.Email.Secret,
		global.Config.Email.Nickname,
		global.Config.Email.Port,
		global.Config.Email.IsSSL,
	))
}
