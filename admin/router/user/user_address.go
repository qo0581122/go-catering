package user

import (
	"catering/api/v1/user_address"

	"github.com/gin-gonic/gin"
)

type UserAddressRouter struct{}

func (d *UserAddressRouter) InitUserAddressRouter(Router *gin.RouterGroup) {
	r := Router.Group("user").Group("addresses").Use()
	{
		r.GET("", user_address.ListPage)
	}
}
