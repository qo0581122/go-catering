package model

type ProductBatch struct {
	BatchName string `json:"batch_name"`
	Status    int    `json:"status"`
	Model
}

func (ProductBatch) TableName() string {
	return "product_batch"
}
