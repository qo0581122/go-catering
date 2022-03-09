package model

type UserIntegrationLog struct {
	Id          uint64 `gorm:"primary_key" json:"id"`
	Uid         int    `json:"uid"`
	ChangeType  int    `json:"change_type"`
	ChangeValue int    `json:"change_value"`
	BeforeValue int    `json:"before_value"`
	AfterValue  int    `json:"after_value"`
	ChangeTime  Time   `json:"change_time"`
}

func (UserIntegrationLog) TableName() string {
	return "user_integration_log"
}
