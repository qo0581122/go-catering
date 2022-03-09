package user

import (
	"catering/api/v1/user_vip_level"
	"catering/api/v1/user_vip_level_log"
	"catering/middleware"

	"github.com/gin-gonic/gin"
)

type UserVipLevelRouter struct{}

func (d *UserVipLevelRouter) InitUserVipLevelRouter(Router *gin.RouterGroup) {
	router := Router.Group("user").Group("vip")
	r := router.Group("levels").Use()
	{
		r.GET("page", user_vip_level.ListPage)
		r.GET("logs", user_vip_level_log.ListPage)
	}

	r = router.Group("level").Use(middleware.OperationRecord())
	{
		r.PUT("", user_vip_level.Adds)
		r.POST("", user_vip_level.Update)
		r.DELETE("", user_vip_level.Delete)
	}
}
