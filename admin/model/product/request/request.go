package request

type CategoryQueryParams struct {
	PageSize     int    `uri:"pageSize" json:"pageSize" form:"pageSize" valid:"Required"`
	PageNum      int    `uri:"pageNum" json:"pageNum" form:"pageNum" valid:"Required"`
	CategoryName string `uri:"category_name"  form:"category_name" json:"category_name"`
	Status       int    `uri:"status" json:"status" form:"status"`
}

type CategoryAddForm struct {
	Status       int    `form:"status" json:"status" valid:"Range(1,2)"`
	CategoryName string `json:"category_name" form:"category_name" valid:"Required"`
}

type CategoryUpdateForm struct {
	Id           uint64 `form:"id" json:"id" valid:"Required"`
	Status       int    `form:"status" json:"status" valid:"Range(1,2)"`
	CategoryName string `json:"category_name" form:"category_name" valid:"Required"`
}

type BatchQueryParams struct {
	PageSize  int    `uri:"pageSize" json:"pageSize" form:"pageSize" valid:"Required"`
	PageNum   int    `uri:"pageNum" json:"pageNum" form:"pageNum" valid:"Required"`
	BatchName string `uri:"batch_name"  form:"batch_name" json:"batch_name"`
	Status    int    `uri:"status" json:"status" form:"status"`
}

type BatchAddForm struct {
	Status    int    `form:"status" json:"status" valid:"Range(1,2)"`
	BatchName string `json:"batch_name" form:"batch_name" valid:"Required"`
}

type BatchUpdateForm struct {
	Id        uint64 `form:"id" json:"id" valid:"Required"`
	Status    int    `form:"status" json:"status" valid:"Range(1,2)"`
	BatchName string `json:"batch_name" form:"batch_name" valid:"Required"`
}

type AttributeQueryParams struct {
	PageSize      int    `uri:"pageSize" json:"pageSize" form:"pageSize" valid:"Required"`
	PageNum       int    `uri:"pageNum" json:"pageNum" form:"pageNum" valid:"Required"`
	AttributeName string `uri:"attribute_name"  form:"attribute_name" json:"attribute_name"`
	Status        int    `uri:"status" json:"status" form:"status"`
}

type AttributeAddForm struct {
	Status          int      `form:"status" json:"status" valid:"Range(1,2)"`
	AttributeName   string   `json:"attribute_name" form:"attribute_name" valid:"Required"`
	AttributeValues []string `json:"values" form:"values"`
}
type AttributeUpdateForm struct {
	Id              uint64   `form:"id" json:"id" valid:"Required"`
	Status          int      `form:"status" json:"status" valid:"Range(1,2)"`
	AttributeName   string   `json:"attribute_name" form:"attribute_name" valid:"Required"`
	AttributeValues []string `json:"values" form:"values"`
}

type AttributeValueQueryParams struct {
	PageSize       int    `uri:"pageSize" json:"pageSize" form:"pageSize" valid:"Required"`
	PageNum        int    `uri:"pageNum" json:"pageNum" form:"pageNum" valid:"Required"`
	AttributeValue string `uri:"attribute_value"  form:"attribute_value" json:"attribute_value"`
	AttributeId    uint64 `uri:"attribute_id" form:"attribute_id" json:"attribute_id"`
	Status         int    `uri:"status" json:"status" form:"status"`
}

type AttributeValueAddForm struct {
	Status         int    `form:"status" json:"status" valid:"Range(1,2)"`
	AttributeValue string `json:"attribute_value" form:"attribute_value" valid:"Required"`
	AttributeId    uint64 `form:"attribute_id" json:"attribute_id" valid:"Required"`
}
type AttributeValueUpdateForm struct {
	Id             uint64 `form:"id" json:"id" valid:"Required"`
	Status         int    `form:"status" json:"status" valid:"Range(1,2)"`
	AttributeValue string `json:"attribute_value" form:"attribute_value" valid:"Required"`
	AttributeId    uint64 `form:"attribute_id" json:"attribute_id" valid:"Required"`
}

type ProductQueryParams struct {
	PageSize    int    `uri:"pageSize" json:"pageSize" form:"pageSize" valid:"Required"`
	PageNum     int    `uri:"pageNum" json:"pageNum" form:"pageNum" valid:"Required"`
	ProductName string `uri:"product_name"  form:"product_name" json:"product_name"`
	Status      int    `uri:"status" json:"status" form:"status"`
}

type ProductAddForm struct {
	ProductName   string   `json:"product_name" form:"product_name" valid:"Required"`
	Description   string   `json:"description" form:"description" valid:"Required"`
	Specis        int      `json:"specis" form:"specis" valid:"Range(1,2)"`
	Url           string   `json:"url" form:"url" valid:"Required"`
	Status        int      `form:"status" json:"status" valid:"Range(1,2)"`
	CategoryId    uint64   `json:"category_id" form:"category_id"`
	AttributeIds  []uint64 `json:"attribute_ids" form:"attribute_ids"`
	BatchIds      []uint64 `json:"batch_ids" form:"batch_ids"`
	ProductIds    []uint64 `json:"product_ids" form:"product_ids" `
	Discount      int      `json:"discount" form:"discount"`
	OriginalPrice int      `json:"original_price" form:"original_price"`
	PayPrice      int      `json:"pay_price" form:"pay_price"`
}

type ProductUpdateForm struct {
	Id           uint64   `form:"id" json:"id" valid:"Required"`
	ProductName  string   `json:"product_name" form:"product_name" valid:"Required"`
	Description  string   `json:"description" form:"description" valid:"Required"`
	Specis       int      `json:"specis" form:"specis" valid:"Range(1,2)"`
	Url          string   `json:"url" form:"url" valid:"Required"`
	Status       int      `form:"status" json:"status" valid:"Range(1,2)"`
	CategoryId   uint64   `json:"category_id" form:"category_id"`
	AttributeIds []uint64 `json:"attribute_ids" form:"attribute_ids"`
	BatchIds     []uint64 `json:"batch_ids" form:"batch_ids"`
	ProductIds   []uint64 `json:"product_ids" form:"product_ids" `
}
