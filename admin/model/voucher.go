package model

type Voucher struct {
	Model
	Name          string `gorm:"column:voucher_name" json:"voucher_name"`
	LeastUsePrice uint64 `json:"least_use_price"`
	Price         uint64 `json:"price"`
	Description   string `json:"description"`
	TotalCount    uint64 `json:"total_count"`
	RemainCount   uint64 `json:"remain_count"`
	GetType       uint32 `json:"get_type"`
	ValidTimeType uint32 `json:"valid_time_type"`
	ValidDay      uint32 `json:"valid_day"`
	UseType       uint32 `json:"use_type"`
	Status        uint32 `json:"status"`
	GetBeginTime  Time   `json:"get_begin_time"`
	GetEndTime    Time   `json:"get_end_time"`
	UseBeginTime  Time   `json:"use_begin_time"`
	UseEndTime    Time   `json:"use_end_time"`
	GetCount      uint32 `json:"get_count"`
}

func (*Voucher) TableName() string {
	return "voucher"
}
