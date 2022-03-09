package model

type UserVipLevel struct {
	Id                uint64 `gorm:"primary_key" json:"id"`
	LevelName         string `json:"level_name"`
	Level             int    `json:"level"`
	UpNeedIntegration int    `json:"up_need_integration"`
	LevelDiscount     int    `json:"level_discount"`
	NextLevelId       uint64 `json:"next_level_id"`
	CreatedAt         Time   `gorm:"column:created_time" json:"created_time"`
	UpdatedAt         Time   `gorm:"column:updated_time" json:"updated_time"`
}

func (UserVipLevel) TableName() string {
	return "user_vip_level"
}
