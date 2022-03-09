package user

import (
	"catering/api/v1/user_address_tag"
	"catering/middleware"

	"github.com/gin-gonic/gin"
)

type UserAddressTagRouter struct{}

func (d *UserAddressTagRouter) InitUserAddressTagRouter(Router *gin.RouterGroup) {
	router := Router.Group("user").Group("address")
	r := router.Group("tags").Use()
	{
		r.GET("", user_address_tag.ListPage)
	}
	r = router.Group("tag").Use(middleware.OperationRecord())
	{
		r.PUT("", user_address_tag.Add)
		r.POST("", user_address_tag.Update)
		r.DELETE(":id", user_address_tag.Delete)
	}
}
