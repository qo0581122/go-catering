package order

import (
	"catering/global"
	"catering/model"
	"catering/model/common/response"
	"catering/model/order"
	"catering/model/order/request"

	"gorm.io/gorm/clause"
)

var OrderService = NewOrderService()

func NewOrderService() orderService {
	return orderServiceImpl{}
}

type orderServiceImpl struct {
}

func (orderServiceImpl) ListPage(pageNum, pageSize int, params *request.RequestParams) *response.ApiResponse {
	var result []order.Order
	o := order.Order{
		Id:          params.OrderId,
		Uid:         params.Uid,
		OrderType:   params.OrderType,
		OrderNumber: params.OrderNumber,
		OrderStatus: params.OrderStatus,
	}
	global.DB.Preload("OrderShop.Shop").Preload(clause.Associations).Scopes(model.Paginate(pageNum, pageSize)).Find(&result)
	var total int64
	global.DB.Model(&o).Count(&total)
	return &response.ApiResponse{
		List:  result,
		Total: int(total),
	}
}

func (orderServiceImpl) GetById(id uint64) *order.Order {
	return nil
}
func (orderServiceImpl) List(params *request.RequestParams) []*order.Order {
	return []*order.Order{}
}
func (orderServiceImpl) Count() int {
	return 0
}
