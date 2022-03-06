package user

import (
	"catering/api/v1/user_address_tag"
	"catering/middleware"

	"github.com/gin-gonic/gin"
)

type UserAddressTagRouter struct{}

func (d *UserAddressTagRouter) InitUserAddressTagRouter(Router *gin.RouterGroup) {
	r := Router.Group("user-address-tags").Use(middleware.OperationRecord())
	{
		r.GET("", user_address_tag.List)
	}
	r = Router.Group("user-address-tag").Use(middleware.OperationRecord())
	{
		r.POST("create", user_address_tag.Add)
		r.POST("update", user_address_tag.Update)
		r.DELETE("delete/:id", user_address_tag.Delete)
	}
}
