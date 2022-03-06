package model

type VoucherGetLog struct {
	Id           uint64  `gorm:"primary_key" json:"id"`
	Uid          int     `json:"uid"`
	VoucherId    uint64  `json:"voucher_id"`
	GetTime      Time    `json:"get_time"`
	UseBeginTime Time    `json:"use_begin_time"`
	UseEndTime   Time    `json:"use_end_time"`
	UseStatus    uint32  `json:"use_status"`
	GetType      uint32  `json:"get_type"`
	Voucher      Voucher `gorm:"foreignKey:VoucherId" json:"voucher"`
}

func (*VoucherGetLog) TableName() string {
	return "user_voucher"
}
