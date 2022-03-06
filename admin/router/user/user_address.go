package user

import (
	"catering/api/v1/user_address"
	"catering/middleware"

	"github.com/gin-gonic/gin"
)

type UserAddressRouter struct{}

func (d *UserAddressRouter) InitUserAddressRouter(Router *gin.RouterGroup) {
	r := Router.Group("user-address").Use(middleware.OperationRecord())
	{
		r.GET("", user_address.ListUserAddress)
	}
}
