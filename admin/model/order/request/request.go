package request

type RequestParams struct {
	OrderId     uint64 `uri:"order_id" json:"order_id" form:"order_id"`
	Uid         uint64 `uri:"uid" json:"uid" form:"uid"`
	OrderType   uint32 `uri:"order_type" json:"order_type" form:"order_type"`
	OrderNumber string `uri:"order_number" json:"order_number" form:"order_number"`
	OrderStatus uint32 `uri:"order_status" json:"order_status" form:"order_status"`
	ShopId      uint64 `uri:"shop_id" json:"shop_id" form:"shop_id"`
}

type ProductIds struct {
	Ids string `uri:"product_ids" json:"product_ids" form:"product_ids"`
}
