package response

import "catering/model"

type ResponseData struct {
	Product      *model.Product   `json:"product"`
	ChildProduct []*model.Product `json:"child_product"`
}
