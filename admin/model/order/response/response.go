package response

import "catering/model"

type OrderProductDetail struct {
	model.Product
	CategoryName string `json:"categroy_name"`
}
