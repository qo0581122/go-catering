package user

import (
	"catering/api/v1/user_vip_level"
	"catering/api/v1/user_vip_level_log"
	"catering/middleware"

	"github.com/gin-gonic/gin"
)

type UserVipLevelRouter struct{}

func (d *UserVipLevelRouter) InitUserVipLevelRouter(Router *gin.RouterGroup) {
	r := Router.Group("user-vip-levels").Use(middleware.OperationRecord())
	{
		r.GET("", user_vip_level.List)
		r.GET("logs", user_vip_level_log.ListUserVipLevelLog)
		r.POST("create", user_vip_level.Adds)
	}

	r = Router.Group("user-vip-level").Use(middleware.OperationRecord())
	{
		r.POST("create", user_vip_level.Add)
		r.POST("update", user_vip_level.Update)
		r.DELETE("delete", user_vip_level.Delete)
	}
}
