package router

import (
	"catering/router/area"
	"catering/router/coupon"
	"catering/router/example"
	"catering/router/order"
	"catering/router/product"
	"catering/router/shop"
	"catering/router/system"
	"catering/router/user"
	"catering/router/voucher"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	Area    area.RouterGroup
	Coupon  coupon.RouterGroup
	Product product.RouterGroup
	Shop    shop.RouterGroup
	User    user.RouterGroup
	Voucher voucher.RouterGroup
	Order   order.OrderRouter
}

var RouterGroupApp = new(RouterGroup)
