package order

import (
	"catering/model/common/response"
	"catering/model/order"
	"catering/model/order/request"
)

type orderService interface {
	ListPage(pageNum, pageSize int, params *request.RequestParams) *response.ApiResponse
	GetById(id uint64) *order.Order
	List(params *request.RequestParams) []*order.Order
	Count() int
}
