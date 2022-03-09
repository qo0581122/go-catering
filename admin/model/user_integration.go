package model

type UserIntegration struct {
	Id                 uint64 `gorm:"primary_key" json:"id"`
	Uid                int    `json:"uid"`
	CurrentIntegration int    `json:"current_integration"`
	HistoryIntegration int    `json:"history_integration"`
	ConsumeIntegration int    `json:"consume_integration"`
	LevelId            uint64 `json:"level_id"`
	CreatedTime        Time   `json:"created_time"`
	UpdatedTime        Time   `json:"updated_time"`
}

func (UserIntegration) TableName() string {
	return "user_integration"
}
